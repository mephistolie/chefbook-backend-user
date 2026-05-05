package user

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-user/internal/entity"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGetUserInfoAddsAvatarLink(t *testing.T) {
	ctx := context.Background()
	userId := uuid.New()
	avatarId := uuid.New()
	firstName := "Test"
	repo := &fakeUserRepo{
		info: entity.UserInfo{
			UserId:    userId,
			FirstName: &firstName,
			AvatarId:  &avatarId,
		},
	}
	s3 := &fakeS3Repo{avatarLinks: map[uuid.UUID]string{avatarId: "https://cdn/avatar.png"}}
	service := NewService(repo, s3)

	info, err := service.GetUserInfo(ctx, userId)
	if err != nil {
		t.Fatalf("GetUserInfo returned error: %v", err)
	}
	if info.AvatarLink == nil || *info.AvatarLink != "https://cdn/avatar.png" {
		t.Fatalf("expected avatar link to be enriched, got %#v", info.AvatarLink)
	}
	if s3.getLinkCalls != 1 {
		t.Fatalf("expected one avatar link lookup, got %d", s3.getLinkCalls)
	}
}

func TestConfirmUserAvatarUploadingRejectsInvalidLink(t *testing.T) {
	ctx := context.Background()
	userId := uuid.New()
	repo := &fakeUserRepo{}
	s3 := &fakeS3Repo{}
	service := NewService(repo, s3)

	err := service.ConfirmUserAvatarUploading(ctx, userId, "not-owned-link")
	if status.Code(err) != codes.InvalidArgument {
		t.Fatalf("expected invalid body, got %s: %v", status.Code(err), err)
	}
	if s3.checkExistsCalls != 0 {
		t.Fatalf("expected no object existence check for invalid link, got %d", s3.checkExistsCalls)
	}
	if repo.setAvatarCalls != 0 {
		t.Fatalf("expected no avatar update for invalid link, got %d", repo.setAvatarCalls)
	}
}

func TestConfirmUserAvatarUploadingRejectsMissingObject(t *testing.T) {
	ctx := context.Background()
	userId := uuid.New()
	avatarId := uuid.New()
	repo := &fakeUserRepo{}
	s3 := &fakeS3Repo{
		avatarIdsByLink: map[string]uuid.UUID{"valid-link": avatarId},
		avatarsExist:    map[uuid.UUID]bool{avatarId: false},
	}
	service := NewService(repo, s3)

	err := service.ConfirmUserAvatarUploading(ctx, userId, "valid-link")
	if status.Code(err) != codes.InvalidArgument || status.Convert(err).Message() != "not found" {
		t.Fatalf("expected missing object error, got %s %q: %v", status.Code(err), status.Convert(err).Message(), err)
	}
	if repo.setAvatarCalls != 0 {
		t.Fatalf("expected no avatar update for missing object, got %d", repo.setAvatarCalls)
	}
}

func TestConfirmUserAvatarUploadingUpdatesAvatarAndDeletesPreviousOne(t *testing.T) {
	ctx := context.Background()
	userId := uuid.New()
	newAvatarId := uuid.New()
	previousAvatarId := uuid.New()
	repo := &fakeUserRepo{previousAvatarId: &previousAvatarId}
	s3 := &fakeS3Repo{
		avatarIdsByLink: map[string]uuid.UUID{"valid-link": newAvatarId},
		avatarsExist:    map[uuid.UUID]bool{newAvatarId: true},
		deleted:         make(chan uuid.UUID, 1),
	}
	service := NewService(repo, s3)

	if err := service.ConfirmUserAvatarUploading(ctx, userId, "valid-link"); err != nil {
		t.Fatalf("ConfirmUserAvatarUploading returned error: %v", err)
	}
	if repo.setAvatarCalls != 1 {
		t.Fatalf("expected one avatar update, got %d", repo.setAvatarCalls)
	}
	if repo.lastAvatarId == nil || *repo.lastAvatarId != newAvatarId {
		t.Fatalf("expected avatar %s to be saved, got %#v", newAvatarId, repo.lastAvatarId)
	}

	select {
	case deletedAvatarId := <-s3.deleted:
		if deletedAvatarId != previousAvatarId {
			t.Fatalf("expected previous avatar %s to be deleted, got %s", previousAvatarId, deletedAvatarId)
		}
	case <-time.After(time.Second):
		t.Fatal("expected previous avatar to be deleted asynchronously")
	}
}

func TestDeleteUserAvatarClearsAvatarAndDeletesPreviousOne(t *testing.T) {
	ctx := context.Background()
	userId := uuid.New()
	previousAvatarId := uuid.New()
	repo := &fakeUserRepo{previousAvatarId: &previousAvatarId}
	s3 := &fakeS3Repo{deleted: make(chan uuid.UUID, 1)}
	service := NewService(repo, s3)

	if err := service.DeleteUserAvatar(ctx, userId); err != nil {
		t.Fatalf("DeleteUserAvatar returned error: %v", err)
	}
	if repo.setAvatarCalls != 1 {
		t.Fatalf("expected one avatar update, got %d", repo.setAvatarCalls)
	}
	if repo.lastAvatarId != nil {
		t.Fatalf("expected avatar to be cleared, got %#v", repo.lastAvatarId)
	}

	select {
	case deletedAvatarId := <-s3.deleted:
		if deletedAvatarId != previousAvatarId {
			t.Fatalf("expected previous avatar %s to be deleted, got %s", previousAvatarId, deletedAvatarId)
		}
	case <-time.After(time.Second):
		t.Fatal("expected previous avatar to be deleted asynchronously")
	}
}

type fakeUserRepo struct {
	info             entity.UserInfo
	previousAvatarId *uuid.UUID
	setAvatarCalls   int
	lastAvatarId     *uuid.UUID
}

func (r *fakeUserRepo) CreateUser(context.Context, uuid.UUID, uuid.UUID) error {
	return errors.New("not implemented")
}

func (r *fakeUserRepo) ImportFirebaseName(context.Context, uuid.UUID, *string, uuid.UUID) error {
	return errors.New("not implemented")
}

func (r *fakeUserRepo) DeleteUser(context.Context, uuid.UUID, uuid.UUID) error {
	return errors.New("not implemented")
}

func (r *fakeUserRepo) GetUsersMinimalInfos(context.Context, []uuid.UUID) map[uuid.UUID]entity.UserMinimalInfo {
	return nil
}

func (r *fakeUserRepo) GetUserInfo(context.Context, uuid.UUID) (entity.UserInfo, error) {
	return r.info, nil
}

func (r *fakeUserRepo) SetUserName(context.Context, uuid.UUID, *string, *string) error {
	return errors.New("not implemented")
}

func (r *fakeUserRepo) SetUserDescription(context.Context, uuid.UUID, *string) error {
	return errors.New("not implemented")
}

func (r *fakeUserRepo) RegisterAvatarUploading(context.Context, uuid.UUID) (uuid.UUID, error) {
	return uuid.Nil, errors.New("not implemented")
}

func (r *fakeUserRepo) SetUserAvatar(_ context.Context, _ uuid.UUID, avatarId *uuid.UUID) (*uuid.UUID, error) {
	r.setAvatarCalls += 1
	r.lastAvatarId = avatarId
	return r.previousAvatarId, nil
}

type fakeS3Repo struct {
	avatarLinks      map[uuid.UUID]string
	avatarIdsByLink  map[string]uuid.UUID
	avatarsExist     map[uuid.UUID]bool
	deleted          chan uuid.UUID
	getLinkCalls     int
	checkExistsCalls int
}

func (r *fakeS3Repo) GetUserAvatarLink(_, avatarId uuid.UUID) string {
	r.getLinkCalls += 1
	return r.avatarLinks[avatarId]
}

func (r *fakeS3Repo) GenerateUserAvatarUploadLink(context.Context, uuid.UUID, uuid.UUID) (entity.PictureUpload, error) {
	return entity.PictureUpload{}, errors.New("not implemented")
}

func (r *fakeS3Repo) CheckAvatarExists(_ context.Context, _ uuid.UUID, avatarId uuid.UUID) bool {
	r.checkExistsCalls += 1
	return r.avatarsExist[avatarId]
}

func (r *fakeS3Repo) DeleteAvatar(_ context.Context, _ uuid.UUID, avatarId uuid.UUID) error {
	if r.deleted != nil {
		r.deleted <- avatarId
	}
	return nil
}

func (r *fakeS3Repo) GetAvatarIdByLink(_ uuid.UUID, link string) *uuid.UUID {
	if avatarId, ok := r.avatarIdsByLink[link]; ok {
		return &avatarId
	}
	return nil
}
