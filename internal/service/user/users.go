package user

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"github.com/mephistolie/chefbook-backend-user/internal/entity"
)

func (s *Service) GetUserInfo(userId uuid.UUID) (entity.UserInfo, error) {
	return s.repo.GetUserInfo(userId)
}

func (s *Service) SetUserName(userId uuid.UUID, firstName *string, lastName *string) error {
	return s.repo.SetUserName(userId, firstName, lastName)
}

func (s *Service) SetUserDescription(userId uuid.UUID, description *string) error {
	return s.repo.SetUserDescription(userId, description)
}

func (s *Service) GenerateUserAvatarUploadLink(userId uuid.UUID) (string, error) {
	return s.s3.GenerateUserAvatarUploadLink(userId)
}

func (s *Service) ConfirmUserAvatarUploading(userId uuid.UUID) error {
	link := s.s3.GetUserAvatarLink(userId)
	if link == nil {
		return fail.GrpcNotFound
	}
	return s.repo.SetUserAvatar(userId, link)
}

func (s *Service) DeleteUserAvatar(userId uuid.UUID) error {
	if err := s.repo.SetUserAvatar(userId, nil); err != nil {
		return err
	}

	go s.s3.DeleteAvatar(userId)

	return nil
}
