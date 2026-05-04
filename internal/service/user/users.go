package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"github.com/mephistolie/chefbook-backend-user/internal/entity"
)

func (s *Service) GetUsersMinimalInfos(ctx context.Context, userIds []uuid.UUID) map[uuid.UUID]entity.UserMinimalInfo {
	infos := s.repo.GetUsersMinimalInfos(ctx, userIds)
	for id, info := range infos {
		if info.AvatarId != nil {
			link := s.s3.GetUserAvatarLink(info.UserId, *info.AvatarId)
			info.AvatarLink = &link
			infos[id] = info
		}
	}
	return infos
}

func (s *Service) GetUserInfo(ctx context.Context, userId uuid.UUID) (entity.UserInfo, error) {
	info, err := s.repo.GetUserInfo(ctx, userId)
	if err != nil {
		return entity.UserInfo{}, err
	}
	if info.AvatarId != nil {
		link := s.s3.GetUserAvatarLink(info.UserId, *info.AvatarId)
		info.AvatarLink = &link
	}
	return info, nil
}

func (s *Service) SetUserName(ctx context.Context, userId uuid.UUID, firstName *string, lastName *string) error {
	return s.repo.SetUserName(ctx, userId, firstName, lastName)
}

func (s *Service) SetUserDescription(ctx context.Context, userId uuid.UUID, description *string) error {
	return s.repo.SetUserDescription(ctx, userId, description)
}

func (s *Service) GenerateUserAvatarUploadLink(ctx context.Context, userId uuid.UUID) (entity.PictureUpload, error) {
	avatarId, err := s.repo.RegisterAvatarUploading(ctx, userId)
	if err != nil {
		return entity.PictureUpload{}, err
	}

	uploading, err := s.s3.GenerateUserAvatarUploadLink(ctx, userId, avatarId)
	if err != nil {
		return entity.PictureUpload{}, err
	}

	return uploading, nil
}

func (s *Service) ConfirmUserAvatarUploading(ctx context.Context, userId uuid.UUID, avatarLink string) error {
	avatarId := s.s3.GetAvatarIdByLink(userId, avatarLink)
	if avatarId == nil {
		return fail.GrpcInvalidBody
	}

	if !s.s3.CheckAvatarExists(ctx, userId, *avatarId) {
		return fail.GrpcNotFound
	}
	return s.setUserAvatar(ctx, userId, avatarId)
}

func (s *Service) DeleteUserAvatar(ctx context.Context, userId uuid.UUID) error {
	return s.setUserAvatar(ctx, userId, nil)
}

func (s *Service) setUserAvatar(ctx context.Context, userId uuid.UUID, avatarId *uuid.UUID) error {
	previousAvatarId, err := s.repo.SetUserAvatar(ctx, userId, avatarId)
	if err != nil {
		return err
	}

	go func() {
		if previousAvatarId != nil {
			_ = s.s3.DeleteAvatar(context.WithoutCancel(ctx), userId, *previousAvatarId)
		}
	}()

	return nil
}
