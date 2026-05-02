package service

import (
	"context"

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
	GetUsersMinimalInfos(ctx context.Context, userIds []uuid.UUID) map[uuid.UUID]entity.UserMinimalInfo
	GetUserInfo(ctx context.Context, userId uuid.UUID) (entity.UserInfo, error)
	SetUserName(ctx context.Context, userId uuid.UUID, firstName, lastName *string) error
	SetUserDescription(ctx context.Context, userId uuid.UUID, description *string) error
	GenerateUserAvatarUploadLink(ctx context.Context, userId uuid.UUID) (entity.PictureUpload, error)
	ConfirmUserAvatarUploading(ctx context.Context, userId uuid.UUID, avatarLnk string) error
	DeleteUserAvatar(ctx context.Context, userId uuid.UUID) error
}

type MQ interface {
	CreateUser(ctx context.Context, userId uuid.UUID, messageId uuid.UUID) error
	ImportFirebaseProfile(ctx context.Context, userId uuid.UUID, firebaseId string, messageId uuid.UUID) error
	DeleteUser(ctx context.Context, userId uuid.UUID, messageId uuid.UUID) error
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
