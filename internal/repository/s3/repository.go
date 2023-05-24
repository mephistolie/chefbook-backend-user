package s3

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"github.com/mephistolie/chefbook-backend-user/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"time"
)

const (
	usersDir   = "users"
	avatarsDir = "avatars"

	avatarMaxSize = 1024 * 512 // 512KB
)

type Repository struct {
	client *minio.Client
	bucket string
}

func NewRepository(cfg config.S3) (*Repository, error) {
	client, err := minio.New(*cfg.Host, &minio.Options{
		Creds:  credentials.NewStaticV4(*cfg.AccessKeyId, *cfg.SecretAccessKey, ""),
		Secure: true,
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

func (r *Repository) GenerateUserAvatarUploadLink(userId, avatarId uuid.UUID) (string, map[string]string, error) {
	return r.generateImageUploadLink(r.getUserAvatarObjectPath(userId, avatarId))
}

func (r *Repository) DeleteAvatar(userId, avatarId uuid.UUID) error {
	object := r.getUserAvatarObjectPath(userId, avatarId)
	opts := minio.RemoveObjectOptions{ForceDelete: true}
	if err := r.client.RemoveObject(context.Background(), object, r.bucket, opts); err != nil {
		log.Warnf("unable to delete user %s avatar: %s", userId, err)
		return fail.GrpcUnknown
	}
	return nil
}

func (r *Repository) getUserAvatarObjectPath(userId, avatarId uuid.UUID) string {
	return fmt.Sprintf("%s/%s/%s/%s", usersDir, userId, avatarsDir, avatarId)
}

func (r *Repository) generateImageUploadLink(objectName string) (string, map[string]string, error) {
	policy := minio.NewPostPolicy()

	if err := policy.SetBucket(r.bucket); err != nil {
		log.Error("unable to set bucket in post policy: ", err)
		return "", map[string]string{}, fail.GrpcUnknown
	}
	if err := policy.SetKey(objectName); err != nil {
		log.Errorf("unable to set object %s in post policy: %s", objectName, err)
		return "", map[string]string{}, fail.GrpcUnknown
	}
	if err := policy.SetExpires(time.Now().Add(1 * time.Hour)); err != nil {
		log.Errorf("unable to set expiration in post policy: %s", objectName, err)
		return "", map[string]string{}, fail.GrpcUnknown
	}

	presignedUrl, formData, err := r.client.PresignedPostPolicy(context.Background(), policy)
	if err != nil {
		log.Errorf("unable to generate presigned link for uploading object %s: %s", objectName, err)
		return "", map[string]string{}, fail.GrpcUnknown
	}

	return presignedUrl.String(), formData, nil
}
