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

func (s *Service) CreateUser(userId uuid.UUID, messageId uuid.UUID) error {
	return s.repo.CreateUser(userId, messageId)
}

func (s *Service) ImportFirebaseProfile(userId uuid.UUID, firebaseId string, messageId uuid.UUID) error {
	if s.firebase == nil {
		log.Warnf("try to import firebase profile with firebase import disabled")
		return errors.New("firebase import disabled")
	}

	firebaseProfile, err := s.firebase.GetProfile(context.Background(), firebaseId)
	if err != nil {
		log.Warnf("unable to get firebase profile for user %s: %s", userId, err)
		return err
	}

	return s.repo.ImportFirebaseName(userId, firebaseProfile.Username, messageId)
}

func (s *Service) DeleteUser(userId uuid.UUID, messageId uuid.UUID) error {
	return s.repo.DeleteUser(userId, messageId)
}
