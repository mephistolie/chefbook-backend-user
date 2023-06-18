package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"github.com/mephistolie/chefbook-backend-user/internal/entity"
)

func (r *Repository) GetUsersMinimalInfos(userIds []uuid.UUID) map[uuid.UUID]entity.UserMinimalInfo {
	var infos map[uuid.UUID]entity.UserMinimalInfo

	query := fmt.Sprintf(`
		SELECT user_id, first_name, last_name, avatar_id
		FROM %s
		WHERE user_id=ANY($1)
	`, usersTable)

	rows, err := r.db.Query(query, userIds)
	if err != nil {
		log.Error("unable to get minimal info for users: ", err)
		return map[uuid.UUID]entity.UserMinimalInfo{}
	}

	for rows.Next() {
		var info entity.UserMinimalInfo
		var firstName *string
		var lastName *string

		if err = rows.Scan(&info.UserId, &firstName, &lastName, &info.AvatarId); err != nil {
			log.Error("unable to parse minimal info for user: ", err)
			continue
		}

		info.FullName = r.getFullName(firstName, lastName)

		infos[info.UserId] = info
	}

	return infos
}

func (r *Repository) getFullName(firstName, lastName *string) *string {
	fullName := ""
	if firstName != nil {
		fullName += *firstName
	}
	if lastName != nil {
		if firstName != nil {
			fullName += " "
		}
		fullName += *lastName
	}

	if len(fullName) > 0 {
		return &fullName
	}
	return nil
}

func (r *Repository) GetUserInfo(userId uuid.UUID) (entity.UserInfo, error) {
	info := entity.UserInfo{}

	query := fmt.Sprintf(`
		SELECT user_id, first_name, last_name, description, avatar_id
		FROM %s
		WHERE user_id=$1
	`, usersTable)

	row := r.db.QueryRow(query, userId)
	if err := row.Scan(&info.UserId, &info.FirstName, &info.LastName, &info.Description, &info.AvatarId); err != nil {
		log.Warnf("unable to get user %s info: %s", userId, err)
		return entity.UserInfo{}, fail.GrpcNotFound
	}

	return info, nil
}

func (r *Repository) SetUserName(userId uuid.UUID, firstName, lastName *string) error {
	query := fmt.Sprintf(`
		UPDATE %s
		SET first_name=$1, last_name=$2
		WHERE user_id=$3
	`, usersTable)

	if _, err := r.db.Exec(query, firstName, lastName, userId); err != nil {
		log.Warnf("unable to set user %s name: %s", userId, err)
		return fail.GrpcUnknown
	}

	return nil
}

func (r *Repository) SetUserDescription(userId uuid.UUID, description *string) error {
	query := fmt.Sprintf(`
		UPDATE %s
		SET description=$1
		WHERE user_id=$2
	`, usersTable)

	if _, err := r.db.Exec(query, description, userId); err != nil {
		log.Warnf("unable to set user %s description: %s", userId, err)
		return fail.GrpcUnknown
	}

	return nil
}

func (r *Repository) RegisterAvatarUploading(userId uuid.UUID) (uuid.UUID, error) {
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

	if err := r.db.Get(&avatarId, query, userId); err != nil {
		log.Errorf("unable to register avatar uploading for user %s: %s", userId, err)
		return uuid.UUID{}, fail.GrpcUnknown
	}

	return avatarId, nil
}

func (r *Repository) SetUserAvatar(userId uuid.UUID, avatarId *uuid.UUID) (*uuid.UUID, error) {
	var previousAvatarId *uuid.UUID = nil

	getPreviousAvatarIdQuery := fmt.Sprintf(`
		SELECT avatar_id
		FROM %s
		WHERE user_id=$1
	`, usersTable)

	if err := r.db.QueryRow(getPreviousAvatarIdQuery, userId).Scan(&previousAvatarId); err != nil {
		log.Warnf("unable to get user %s avatar id: %s", userId, err)
		return nil, fail.GrpcUnknown
	}

	if (avatarId != nil && previousAvatarId != nil && *avatarId == *previousAvatarId) || avatarId == previousAvatarId {
		return nil, nil
	}

	tx, err := r.startTransaction()
	if err != nil {
		return nil, err
	}

	setAvatarQuery := fmt.Sprintf(`
		UPDATE %s
		SET avatar_id=$1
		WHERE user_id=$2
	`, usersTable)

	if _, err := tx.Exec(setAvatarQuery, avatarId, userId); err != nil {
		log.Warnf("unable to set user %s avatar id: %s", userId, err)
		return nil, fail.GrpcUnknown
	}

	if avatarId != nil {
		deleteUploadingQuery := fmt.Sprintf(`
		DELETE FROM %s
		WHERE avatar_id=$1 AND user_id=$2
	`, avatarUploadsTable)

		if _, err := tx.Exec(deleteUploadingQuery, *avatarId, userId); err != nil {
			log.Warnf("unable to delete avatar uploading record: %s", userId, err)
			return nil, fail.GrpcUnknown
		}
	}

	return previousAvatarId, commitTransaction(tx)
}
