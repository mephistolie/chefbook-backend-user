package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	api "github.com/mephistolie/chefbook-backend-user/api/proto/implementation/v1"
	userFail "github.com/mephistolie/chefbook-backend-user/internal/entity/fail"
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
		dto := api.UserMinInfo{}
		if info.FullName != nil {
			dto.FullName = *info.FullName
		}
		if info.AvatarLink != nil {
			dto.Avatar = *info.AvatarLink
		}

		infos[id.String()] = &dto
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

	firstName, lastName, description, avatar := "", "", "", ""
	if info.FirstName != nil {
		firstName = *info.FirstName
	}
	if info.LastName != nil {
		lastName = *info.LastName
	}
	if info.Description != nil {
		description = *info.Description
	}
	if info.AvatarLink != nil {
		avatar = *info.AvatarLink
	}

	return &api.GetUserInfoResponse{
		UserId:      info.UserId.String(),
		FirstName:   firstName,
		LastName:    lastName,
		Description: description,
		Avatar:      avatar,
	}, nil
}

func (s *UserServer) SetUserName(_ context.Context, req *api.SetUserNameRequest) (*api.SetUserNameResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	if len(req.FirstName) > 64 || len(req.LastName) > 64 {
		return nil, userFail.GrpcNameLength
	}
	var firstName, lastName *string = nil, nil
	if len(req.FirstName) > 0 {
		firstName = &req.FirstName
	}
	if len(req.LastName) > 0 {
		lastName = &req.LastName
	}

	err = s.service.SetUserName(userId, firstName, lastName)
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
	var description *string = nil
	if len(req.Description) > 0 {
		if len(req.Description) > 150 {
			return nil, userFail.GrpcDescriptionLength
		}
		description = &req.Description
	}

	err = s.service.SetUserDescription(userId, description)
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

	avatarId, link, formData, err := s.service.GenerateUserAvatarUploadLink(userId)
	if err != nil {
		return nil, err
	}

	return &api.GenerateUserAvatarUploadLinkResponse{
		AvatarId: avatarId.String(),
		Link:     link,
		FormData: formData,
	}, nil
}

func (s *UserServer) ConfirmUserAvatarUploading(_ context.Context, req *api.ConfirmUserAvatarUploadingRequest) (*api.ConfirmUserAvatarUploadingResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}
	avatarId, err := uuid.Parse(req.AvatarId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	if err = s.service.ConfirmUserAvatarUploading(userId, avatarId); err != nil {
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
