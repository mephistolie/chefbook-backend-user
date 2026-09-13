package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"github.com/mephistolie/chefbook-backend-user/internal/entity"
	"github.com/mephistolie/chefbook-backend-user/internal/logging"
)

func (r *Repository) GetUsersMinimalInfos(ctx context.Context, userIds []uuid.UUID) map[uuid.UUID]entity.UserMinimalInfo {
	infos := make(map[uuid.UUID]entity.UserMinimalInfo)

	query := fmt.Sprintf(`
		SELECT user_id, display_name, avatar_id
		FROM %s
		WHERE user_id=ANY($1)
	`, usersTable)

	rows, err := r.db.QueryContext(ctx, query, userIds)
	if err != nil {
		r.events.MinimalInfoQueryFailed(ctx, logging.BatchOperationData{
			Operation:      "query",
			RequestedCount: len(userIds),
		}, err)
		return map[uuid.UUID]entity.UserMinimalInfo{}
	}
	defer rows.Close()

	for rows.Next() {
		var info entity.UserMinimalInfo

		if err = rows.Scan(&info.UserId, &info.DisplayName, &info.AvatarId); err != nil {
			r.events.MinimalInfoQueryDegraded(ctx, logging.BatchOperationData{
				Operation:      "scan",
				RequestedCount: len(userIds),
			}, err)
			continue
		}

		infos[info.UserId] = info
	}
	if err = rows.Err(); err != nil {
		r.events.MinimalInfoQueryFailed(ctx, logging.BatchOperationData{
			Operation:      "iterate",
			RequestedCount: len(userIds),
		}, err)
		return map[uuid.UUID]entity.UserMinimalInfo{}
	}

	return infos
}

func (r *Repository) GetUserInfo(ctx context.Context, userId uuid.UUID) (entity.UserInfo, error) {
	info := entity.UserInfo{}

	query := fmt.Sprintf(`
		SELECT user_id, display_name, description, avatar_id
		FROM %s
		WHERE user_id=$1
	`, usersTable)

	row := r.db.QueryRowContext(ctx, query, userId)
	if err := row.Scan(&info.UserId, &info.DisplayName, &info.Description, &info.AvatarId); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			r.events.UserReadFailed(ctx, logging.UserOperationData{
				UserID:    userId.String(),
				Operation: "get_info",
			}, err)
		}
		return entity.UserInfo{}, fail.GrpcNotFound
	}

	return info, nil
}

func (r *Repository) SetUserDisplayName(ctx context.Context, userId uuid.UUID, displayName *string) error {
	query := fmt.Sprintf(`
		UPDATE %s
		SET display_name=$1
		WHERE user_id=$2
	`, usersTable)

	if _, err := r.db.ExecContext(ctx, query, displayName, userId); err != nil {
		r.events.UserMutationFailed(ctx, logging.UserOperationData{
			UserID:    userId.String(),
			Operation: "set_name",
		}, err)
		return fail.GrpcUnknown
	}

	return nil
}

func (r *Repository) SetUserDescription(ctx context.Context, userId uuid.UUID, description *string) error {
	query := fmt.Sprintf(`
		UPDATE %s
		SET description=$1
		WHERE user_id=$2
	`, usersTable)

	if _, err := r.db.ExecContext(ctx, query, description, userId); err != nil {
		r.events.UserMutationFailed(ctx, logging.UserOperationData{
			UserID:    userId.String(),
			Operation: "set_description",
		}, err)
		return fail.GrpcUnknown
	}

	return nil
}

func (r *Repository) RegisterAvatarUploading(ctx context.Context, userId uuid.UUID) (uuid.UUID, error) {
	var avatarId uuid.UUID

	query := fmt.Sprintf(`
		WITH s AS
		(
			SELECT avatar_id
			FROM %[1]v
			WHERE user_id=$1
		), i AS
		(
			INSERT INTO %[1]v (user_id)
			SELECT $1
			WHERE NOT EXISTS (SELECT 1 FROM s)
			RETURNING avatar_id
		)
		SELECT avatar_id FROM i
		UNION ALL
		SELECT avatar_id FROM s
	`, avatarUploadsTable)

	if err := r.db.GetContext(ctx, &avatarId, query, userId); err != nil {
		r.events.UserMutationFailed(ctx, logging.UserOperationData{
			UserID:    userId.String(),
			Operation: "register_avatar_upload",
		}, err)
		return uuid.UUID{}, fail.GrpcUnknown
	}

	return avatarId, nil
}

func (r *Repository) SetUserAvatar(ctx context.Context, userId uuid.UUID, avatarId *uuid.UUID) (*uuid.UUID, error) {
	var previousAvatarId *uuid.UUID = nil

	getPreviousAvatarIdQuery := fmt.Sprintf(`
		SELECT avatar_id
		FROM %s
		WHERE user_id=$1
	`, usersTable)

	if err := r.db.QueryRowContext(ctx, getPreviousAvatarIdQuery, userId).Scan(&previousAvatarId); err != nil {
		r.events.UserReadFailed(ctx, logging.UserOperationData{
			UserID:    userId.String(),
			Operation: "read_avatar",
		}, err)
		return nil, fail.GrpcUnknown
	}

	if (avatarId != nil && previousAvatarId != nil && *avatarId == *previousAvatarId) || avatarId == previousAvatarId {
		return nil, nil
	}

	tx, err := r.startTransaction(ctx)
	if err != nil {
		return nil, err
	}

	setAvatarQuery := fmt.Sprintf(`
		UPDATE %s
		SET avatar_id=$1
		WHERE user_id=$2
	`, usersTable)

	if _, err := tx.ExecContext(ctx, setAvatarQuery, avatarId, userId); err != nil {
		r.events.UserMutationFailed(ctx, logging.UserOperationData{
			UserID:    userId.String(),
			Operation: "set_avatar",
		}, err)
		return nil, fail.GrpcUnknown
	}

	if avatarId != nil {
		deleteUploadingQuery := fmt.Sprintf(`
		DELETE FROM %s
		WHERE avatar_id=$1 AND user_id=$2
	`, avatarUploadsTable)

		if _, err := tx.ExecContext(ctx, deleteUploadingQuery, *avatarId, userId); err != nil {
			r.events.UserMutationFailed(ctx, logging.UserOperationData{
				UserID:    userId.String(),
				Operation: "delete_avatar_upload",
			}, err)
			return nil, fail.GrpcUnknown
		}
	}

	return previousAvatarId, r.commitTransaction(ctx, tx)
}
