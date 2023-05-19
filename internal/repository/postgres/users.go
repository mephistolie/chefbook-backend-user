package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"github.com/mephistolie/chefbook-backend-user/internal/entity"
)

func (r *Repository) GetUserInfo(userId uuid.UUID) (entity.UserInfo, error) {
	info := entity.UserInfo{}

	query := fmt.Sprintf(`
			SELECT user_id, first_name, last_name, description, avatar
			FROM %s
			WHERE user_id=$1
		`, usersTable)

	row := r.db.QueryRow(query, userId)
	if err := row.Scan(&info.UserId, &info.FirstName, &info.LastName, &info.Description, &info.AvatarUrl); err != nil {
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

func (r *Repository) SetUserAvatar(userId uuid.UUID, link *string) error {
	query := fmt.Sprintf(`
			UPDATE %s
			SET avatar=$1
			WHERE user_id=$2
		`, usersTable)

	if _, err := r.db.Exec(query, link, userId); err != nil {
		log.Warnf("unable to set user %s avatar link: %s", userId, err)
		return fail.GrpcUnknown
	}

	return nil
}
