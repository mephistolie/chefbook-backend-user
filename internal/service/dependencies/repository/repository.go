package repository

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-user/internal/entity"
)

type User interface {
	CreateUser(userId uuid.UUID, messageId uuid.UUID) error
	ImportFirebaseName(userId uuid.UUID, username *string, messageId uuid.UUID) error
	DeleteUser(userId uuid.UUID, messageId uuid.UUID) error

	GetUserInfo(userId uuid.UUID) (entity.UserInfo, error)
	SetUserName(userId uuid.UUID, firstName, lastName *string) error
	SetUserDescription(userId uuid.UUID, description *string) error
	SetUserAvatar(userId uuid.UUID, link *string) error
}

type S3 interface {
	GetUserAvatarLink(userId uuid.UUID) *string
	GenerateUserAvatarUploadLink(userId uuid.UUID) (string, error)
	DeleteAvatar(userId uuid.UUID) error
}
