package repository

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-user/internal/entity"
)

type User interface {
	CreateUser(userId uuid.UUID, messageId uuid.UUID) error
	ImportFirebaseName(userId uuid.UUID, username *string, messageId uuid.UUID) error
	DeleteUser(userId uuid.UUID, messageId uuid.UUID) error

	GetUsersMinimalInfos(userIds []uuid.UUID) map[uuid.UUID]entity.UserMinimalInfo
	GetUserInfo(userId uuid.UUID) (entity.UserInfo, error)
	SetUserName(userId uuid.UUID, firstName, lastName *string) error
	SetUserDescription(userId uuid.UUID, description *string) error
	RegisterAvatarUploading(userId uuid.UUID) (uuid.UUID, error)
	SetUserAvatar(userId uuid.UUID, avatarId *uuid.UUID) (*uuid.UUID, error)
}

type S3 interface {
	GetUserAvatarLink(userId, avatarId uuid.UUID) string
	GenerateUserAvatarUploadLink(userId, avatarId uuid.UUID) (string, map[string]string, error)
	CheckAvatarExists(userId, avatarId uuid.UUID) bool
	DeleteAvatar(userId, avatarId uuid.UUID) error
}
