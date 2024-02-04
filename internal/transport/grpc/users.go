package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	api "github.com/mephistolie/chefbook-backend-user/api/proto/implementation/v1"
)

const (
	maxNameLength        = 64
	maxDescriptionLength = 150
)

func (s *UserServer) GetUsersMinInfo(_ context.Context, req *api.GetUsersMinInfoRequest) (*api.GetUsersMinInfoResponse, error) {
	var userIds []uuid.UUID
	for _, rawId := range req.UserIds {
		if userId, err := uuid.Parse(rawId); err == nil {
			userIds = append(userIds, userId)
		}
	}

	response := s.service.User.GetUsersMinimalInfos(userIds)

	infos := make(map[string]*api.UserMinInfo)
	for id, info := range response {
		infos[id.String()] = &api.UserMinInfo{
			FullName: info.FullName,
			Avatar:   info.AvatarLink,
		}
	}

	return &api.GetUsersMinInfoResponse{Infos: infos}, nil
}

func (s *UserServer) GetUserInfo(_ context.Context, req *api.GetUserInfoRequest) (*api.GetUserInfoResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	info, err := s.service.GetUserInfo(userId)
	if err != nil {
		return nil, err
	}

	return &api.GetUserInfoResponse{
		UserId:      info.UserId.String(),
		FirstName:   info.FirstName,
		LastName:    info.LastName,
		Description: info.Description,
		Avatar:      info.AvatarLink,
	}, nil
}

func (s *UserServer) SetUserName(_ context.Context, req *api.SetUserNameRequest) (*api.SetUserNameResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	if req.FirstName != nil && len([]rune(*req.FirstName)) > maxNameLength {
		firstName := string([]rune(*req.FirstName)[0:maxNameLength])
		req.FirstName = &firstName
	}
	if req.LastName != nil && len([]rune(*req.LastName)) > maxNameLength {
		lastName := string([]rune(*req.LastName)[0:maxNameLength])
		req.LastName = &lastName
	}

	err = s.service.SetUserName(userId, req.FirstName, req.LastName)
	if err != nil {
		return nil, err
	}

	return &api.SetUserNameResponse{Message: "user name changed"}, nil
}

func (s *UserServer) SetUserDescription(_ context.Context, req *api.SetUserDescriptionRequest) (*api.SetUserDescriptionResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	if req.Description != nil && len([]rune(*req.Description)) > maxDescriptionLength {
		description := string([]rune(*req.Description)[0:maxDescriptionLength])
		req.Description = &description
	}

	err = s.service.SetUserDescription(userId, req.Description)
	if err != nil {
		return nil, err
	}

	return &api.SetUserDescriptionResponse{Message: "user description changed"}, nil
}

func (s *UserServer) GenerateUserAvatarUploadLink(_ context.Context, req *api.GenerateUserAvatarUploadLinkRequest) (*api.GenerateUserAvatarUploadLinkResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	uploading, err := s.service.GenerateUserAvatarUploadLink(userId)
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

func (s *UserServer) ConfirmUserAvatarUploading(_ context.Context, req *api.ConfirmUserAvatarUploadingRequest) (*api.ConfirmUserAvatarUploadingResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	if err = s.service.ConfirmUserAvatarUploading(userId, req.AvatarLink); err != nil {
		return nil, err
	}

	return &api.ConfirmUserAvatarUploadingResponse{Message: "new avatar applied"}, nil
}

func (s *UserServer) DeleteUserAvatar(_ context.Context, req *api.DeleteUserAvatarRequest) (*api.DeleteUserAvatarResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	if err = s.service.DeleteUserAvatar(userId); err != nil {
		return nil, err
	}

	return &api.DeleteUserAvatarResponse{Message: "avatar deleted"}, nil
}
