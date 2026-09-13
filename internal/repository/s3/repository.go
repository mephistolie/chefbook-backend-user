package s3

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"github.com/mephistolie/chefbook-backend-user/internal/config"
	"github.com/mephistolie/chefbook-backend-user/internal/entity"
	"github.com/mephistolie/chefbook-backend-user/internal/logging"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	usersDir   = "users"
	avatarsDir = "avatars"

	avatarMaxSize = 1024 * 512 // 512KB
)

type Repository struct {
	client *minio.Client
	bucket string
	events logging.Events
}

func NewRepository(cfg config.S3) (*Repository, error) {
	client, err := minio.New(*cfg.Host, &minio.Options{
		Creds:  credentials.NewStaticV4(*cfg.AccessKeyId, *cfg.SecretAccessKey, ""),
		Secure: true,
		Region: *cfg.Region,
	})
	if err != nil {
		return nil, err
	}

	return &Repository{
		client: client,
		bucket: *cfg.Bucket,
	}, nil
}

func (r *Repository) GetUserAvatarLink(userId, avatarId uuid.UUID) string {
	objectPath := r.getUserAvatarObjectPath(userId, avatarId)
	return fmt.Sprintf("https://%s/%s", r.bucket, objectPath)
}

func (r *Repository) GenerateUserAvatarUploadLink(ctx context.Context, userId, avatarId uuid.UUID) (entity.PictureUpload, error) {
	return r.generateImageUploadLink(ctx, r.getUserAvatarObjectPath(userId, avatarId))
}

func (r *Repository) CheckAvatarExists(ctx context.Context, userId, avatarId uuid.UUID) bool {
	_, err := r.client.GetObjectACL(ctx, r.bucket, r.getUserAvatarObjectPath(userId, avatarId))
	return err == nil
}

func (r *Repository) DeleteAvatar(ctx context.Context, userId, avatarId uuid.UUID) error {
	object := r.getUserAvatarObjectPath(userId, avatarId)
	opts := minio.RemoveObjectOptions{ForceDelete: true}
	if err := r.client.RemoveObject(ctx, r.bucket, object, opts); err != nil {
		r.events.AvatarObjectDeleteFailed(ctx, userId.String(), err)
		return fail.GrpcUnknown
	}
	return nil
}

func (r *Repository) getUserAvatarObjectPath(userId, avatarId uuid.UUID) string {
	return fmt.Sprintf("%s/%s/%s/%s", usersDir, userId, avatarsDir, avatarId)
}

func (r *Repository) generateImageUploadLink(ctx context.Context, objectName string) (entity.PictureUpload, error) {
	policy := minio.NewPostPolicy()

	if err := policy.SetBucket(r.bucket); err != nil {
		r.events.AvatarUploadPolicyFailed(ctx, "set_bucket", err)
		return entity.PictureUpload{}, fail.GrpcUnknown
	}
	if err := policy.SetKey(objectName); err != nil {
		r.events.AvatarUploadPolicyFailed(ctx, "set_key", err)
		return entity.PictureUpload{}, fail.GrpcUnknown
	}
	if err := policy.SetContentTypeStartsWith("image"); err != nil {
		r.events.AvatarUploadPolicyFailed(ctx, "set_content_type", err)
		return entity.PictureUpload{}, fail.GrpcUnknown
	}
	if err := policy.SetContentLengthRange(0, avatarMaxSize); err != nil {
		r.events.AvatarUploadPolicyFailed(ctx, "set_content_length", err)
		return entity.PictureUpload{}, fail.GrpcUnknown
	}
	if err := policy.SetExpires(time.Now().Add(1 * time.Hour)); err != nil {
		r.events.AvatarUploadPolicyFailed(ctx, "set_expiration", err)
		return entity.PictureUpload{}, fail.GrpcUnknown
	}

	uploadUrl, formData, err := r.client.PresignedPostPolicy(ctx, policy)
	if err != nil {
		r.events.AvatarUploadLinkGenerationFailed(ctx, err)
		return entity.PictureUpload{}, fail.GrpcUnknown
	}

	return entity.PictureUpload{
		PictureLink: fmt.Sprintf("https://%s/%s", r.bucket, objectName),
		UploadUrl:   uploadUrl.String(),
		FormData:    formData,
		MaxSize:     avatarMaxSize,
	}, nil
}

func (r *Repository) GetAvatarIdByLink(ctx context.Context, userId uuid.UUID, link string) *uuid.UUID {
	pictureUrl, err := url.Parse(link)
	if err != nil || pictureUrl.Host != r.bucket {
		return nil
	}
	fragments := strings.Split(pictureUrl.Path, "/")
	if len(fragments) > 1 && fragments[0] == "" {
		fragments = fragments[1:]
	}
	if len(fragments) != 4 ||
		fragments[0] != usersDir ||
		fragments[1] != userId.String() ||
		fragments[2] != avatarsDir {
		r.events.AvatarLinkRejected(ctx, "invalid_path")
		return nil
	}
	avatarId, err := uuid.Parse(fragments[3])
	if err != nil {
		r.events.AvatarLinkRejected(ctx, "invalid_avatar_id")
		return nil
	}
	return &avatarId
}
