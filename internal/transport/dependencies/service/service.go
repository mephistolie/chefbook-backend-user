package service

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/firebase"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-user/internal/config"
	"github.com/mephistolie/chefbook-backend-user/internal/entity"
	s32 "github.com/mephistolie/chefbook-backend-user/internal/repository/s3"
	"github.com/mephistolie/chefbook-backend-user/internal/service/dependencies/repository"
	"github.com/mephistolie/chefbook-backend-user/internal/service/mq"
	"github.com/mephistolie/chefbook-backend-user/internal/service/user"
)

type Service struct {
	User
	MQ
}

type User interface {
	GetUsersMinimalInfos(userIds []uuid.UUID) map[uuid.UUID]entity.UserMinimalInfo
	GetUserInfo(userId uuid.UUID) (entity.UserInfo, error)
	SetUserName(userId uuid.UUID, firstName, lastName *string) error
	SetUserDescription(userId uuid.UUID, description *string) error
	GenerateUserAvatarUploadLink(userId uuid.UUID) (uuid.UUID, string, map[string]string, error)
	ConfirmUserAvatarUploading(userId uuid.UUID, avatarId uuid.UUID) error
	DeleteUserAvatar(userId uuid.UUID) error
}

type MQ interface {
	CreateUser(userId uuid.UUID, messageId uuid.UUID) error
	ImportFirebaseProfile(userId uuid.UUID, firebaseId string, messageId uuid.UUID) error
	DeleteUser(userId uuid.UUID, messageId uuid.UUID) error
}

func New(
	cfg *config.Config,
	repo repository.User,
) (*Service, error) {
	var err error = nil
	var client *firebase.Client = nil
	if len(*cfg.Firebase.Credentials) > 0 {
		credentials := []byte(*cfg.Firebase.Credentials)
		client, err = firebase.NewClient(credentials, "")
		if err != nil {
			return nil, err
		}
		log.Info("Firebase client initialized")
	}

	s3, err := s32.NewRepository(cfg.S3)
	if err != nil {
		return nil, err
	}

	return &Service{
		User: user.NewService(repo, s3),
		MQ:   mq.NewService(repo, client),
	}, nil
}
