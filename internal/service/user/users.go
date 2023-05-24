package user

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-user/internal/entity"
)

func (s *Service) GetUserInfo(userId uuid.UUID) (entity.UserInfo, error) {
	info, err := s.repo.GetUserInfo(userId)
	if err != nil {
		return entity.UserInfo{}, err
	}
	if info.AvatarId != nil {
		link := s.s3.GetUserAvatarLink(info.UserId, *info.AvatarId)
		info.AvatarLink = &link
	}
	return info, nil
}

func (s *Service) SetUserName(userId uuid.UUID, firstName *string, lastName *string) error {
	return s.repo.SetUserName(userId, firstName, lastName)
}

func (s *Service) SetUserDescription(userId uuid.UUID, description *string) error {
	return s.repo.SetUserDescription(userId, description)
}

func (s *Service) GenerateUserAvatarUploadLink(userId uuid.UUID) (string, error) {
	avatarId, err := s.repo.RegisterAvatarUploading(userId)
	if err != nil {
		return "", err
	}
	return s.s3.GenerateUserAvatarUploadLink(userId, avatarId)
}

func (s *Service) ConfirmUserAvatarUploading(userId uuid.UUID, avatarId uuid.UUID) error {
	return s.setUserAvatar(userId, &avatarId)
}

func (s *Service) DeleteUserAvatar(userId uuid.UUID) error {
	return s.setUserAvatar(userId, nil)
}

func (s *Service) setUserAvatar(userId uuid.UUID, avatarId *uuid.UUID) error {
	previousAvatarId, err := s.repo.SetUserAvatar(userId, avatarId)
	if err != nil {
		return err
	}

	go func() {
		if previousAvatarId != nil {
			_ = s.s3.DeleteAvatar(userId, *previousAvatarId)
		}
	}()

	return nil
}
