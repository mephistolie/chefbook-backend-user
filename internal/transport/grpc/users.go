package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	api "github.com/mephistolie/chefbook-backend-user/api/proto/implementation/v1"
)

const (
	maxNameLength        = 128
	maxDescriptionLength = 150
)

func (s *UserServer) GetUsersMinInfo(ctx context.Context, req *api.GetUsersMinInfoRequest) (*api.GetUsersMinInfoResponse, error) {
	var userIds []uuid.UUID
	for _, rawId := range req.UserIds {
		if userId, err := uuid.Parse(rawId); err == nil {
			userIds = append(userIds, userId)
		}
	}

	response := s.service.User.GetUsersMinimalInfos(ctx, userIds)

	infos := make(map[string]*api.UserMinInfo)
	for id, info := range response {
		infos[id.String()] = &api.UserMinInfo{
			DisplayName: info.DisplayName,
			Avatar:      info.AvatarLink,
		}
	}

	return &api.GetUsersMinInfoResponse{Infos: infos}, nil
}

func (s *UserServer) GetUserInfo(ctx context.Context, req *api.GetUserInfoRequest) (*api.GetUserInfoResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	info, err := s.service.GetUserInfo(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &api.GetUserInfoResponse{
		UserId:      info.UserId.String(),
		DisplayName: info.DisplayName,
		Description: info.Description,
		Avatar:      info.AvatarLink,
	}, nil
}

func (s *UserServer) SetUserDisplayName(ctx context.Context, req *api.SetUserDisplayNameRequest) (*api.SetUserDisplayNameResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	if req.DisplayName != nil && len([]rune(*req.DisplayName)) > maxNameLength {
		displayName := string([]rune(*req.DisplayName)[0:maxNameLength])
		req.DisplayName = &displayName
	}

	err = s.service.SetUserDisplayName(ctx, userId, req.DisplayName)
	if err != nil {
		return nil, err
	}

	return &api.SetUserDisplayNameResponse{Message: "user name changed"}, nil
}

func (s *UserServer) SetUserDescription(ctx context.Context, req *api.SetUserDescriptionRequest) (*api.SetUserDescriptionResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	if req.Description != nil && len([]rune(*req.Description)) > maxDescriptionLength {
		description := string([]rune(*req.Description)[0:maxDescriptionLength])
		req.Description = &description
	}

	err = s.service.SetUserDescription(ctx, userId, req.Description)
	if err != nil {
		return nil, err
	}

	return &api.SetUserDescriptionResponse{Message: "user description changed"}, nil
}

func (s *UserServer) GenerateUserAvatarUploadLink(ctx context.Context, req *api.GenerateUserAvatarUploadLinkRequest) (*api.GenerateUserAvatarUploadLinkResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	uploading, err := s.service.GenerateUserAvatarUploadLink(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &api.GenerateUserAvatarUploadLinkResponse{
		AvatarLink: uploading.PictureLink,
		UploadLink: uploading.UploadUrl,
		FormData:   uploading.FormData,
		MaxSize:    uploading.MaxSize,
	}, nil
}

func (s *UserServer) ConfirmUserAvatarUploading(ctx context.Context, req *api.ConfirmUserAvatarUploadingRequest) (*api.ConfirmUserAvatarUploadingResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	if err = s.service.ConfirmUserAvatarUploading(ctx, userId, req.AvatarLink); err != nil {
		return nil, err
	}

	return &api.ConfirmUserAvatarUploadingResponse{Message: "new avatar applied"}, nil
}

func (s *UserServer) DeleteUserAvatar(ctx context.Context, req *api.DeleteUserAvatarRequest) (*api.DeleteUserAvatarResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	if err = s.service.DeleteUserAvatar(ctx, userId); err != nil {
		return nil, err
	}

	return &api.DeleteUserAvatarResponse{Message: "avatar deleted"}, nil
}
