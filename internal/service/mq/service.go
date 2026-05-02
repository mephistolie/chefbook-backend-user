package mq

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/firebase"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-user/internal/service/dependencies/repository"
)

type Service struct {
	repo     repository.User
	firebase *firebase.Client
}

func NewService(
	repo repository.User,
	firebase *firebase.Client,
) *Service {
	return &Service{
		repo:     repo,
		firebase: firebase,
	}
}

func (s *Service) CreateUser(ctx context.Context, userId uuid.UUID, messageId uuid.UUID) error {
	return s.repo.CreateUser(ctx, userId, messageId)
}

func (s *Service) ImportFirebaseProfile(ctx context.Context, userId uuid.UUID, firebaseId string, messageId uuid.UUID) error {
	if s.firebase == nil {
		log.Warnf("try to import firebase profile with firebase import disabled")
		return errors.New("firebase import disabled")
	}

	firebaseProfile, err := s.firebase.GetProfile(ctx, firebaseId)
	if err != nil {
		log.Warnf("unable to get firebase profile for user %s: %s", userId, err)
		return err
	}

	return s.repo.ImportFirebaseName(ctx, userId, firebaseProfile.Username, messageId)
}

func (s *Service) DeleteUser(ctx context.Context, userId uuid.UUID, messageId uuid.UUID) error {
	return s.repo.DeleteUser(ctx, userId, messageId)
}
