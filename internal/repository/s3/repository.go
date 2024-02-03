package s3

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"github.com/mephistolie/chefbook-backend-user/internal/config"
	"github.com/mephistolie/chefbook-backend-user/internal/entity"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"net/url"
	"strings"
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

func (r *Repository) GenerateUserAvatarUploadLink(userId, avatarId uuid.UUID) (entity.PictureUpload, error) {
	return r.generateImageUploadLink(r.getUserAvatarObjectPath(userId, avatarId))
}

func (r *Repository) CheckAvatarExists(userId, avatarId uuid.UUID) bool {
	_, err := r.client.GetObjectACL(context.Background(), r.bucket, r.getUserAvatarObjectPath(userId, avatarId))
	return err == nil
}

func (r *Repository) DeleteAvatar(userId, avatarId uuid.UUID) error {
	object := r.getUserAvatarObjectPath(userId, avatarId)
	opts := minio.RemoveObjectOptions{ForceDelete: true}
	if err := r.client.RemoveObject(context.Background(), r.bucket, object, opts); err != nil {
		log.Warnf("unable to delete user %s avatar: %s", userId, err)
		return fail.GrpcUnknown
	}
	return nil
}

func (r *Repository) getUserAvatarObjectPath(userId, avatarId uuid.UUID) string {
	return fmt.Sprintf("%s/%s/%s/%s", usersDir, userId, avatarsDir, avatarId)
}

func (r *Repository) generateImageUploadLink(objectName string) (entity.PictureUpload, error) {
	policy := minio.NewPostPolicy()

	if err := policy.SetBucket(r.bucket); err != nil {
		log.Error("unable to set bucket in post policy: ", err)
		return entity.PictureUpload{}, fail.GrpcUnknown
	}
	if err := policy.SetKey(objectName); err != nil {
		log.Errorf("unable to set object %s in post policy: %s", objectName, err)
		return entity.PictureUpload{}, fail.GrpcUnknown
	}
	if err := policy.SetContentTypeStartsWith("image"); err != nil {
		log.Errorf("unable to set content type in post policy: %s", objectName, err)
		return entity.PictureUpload{}, fail.GrpcUnknown
	}
	if err := policy.SetContentLengthRange(0, avatarMaxSize); err != nil {
		log.Errorf("unable to set content length in post policy: %s", objectName, err)
		return entity.PictureUpload{}, fail.GrpcUnknown
	}
	if err := policy.SetExpires(time.Now().Add(1 * time.Hour)); err != nil {
		log.Errorf("unable to set expiration in post policy: %s", objectName, err)
		return entity.PictureUpload{}, fail.GrpcUnknown
	}

	uploadUrl, formData, err := r.client.PresignedPostPolicy(context.Background(), policy)
	if err != nil {
		log.Errorf("unable to generate presigned link for uploading object %s: %s", objectName, err)
		return entity.PictureUpload{}, fail.GrpcUnknown
	}

	return entity.PictureUpload{
		PictureLink: fmt.Sprintf("https://%s", r.bucket),
		UploadUrl:   uploadUrl.String(),
		FormData:    formData,
		MaxSize:     avatarMaxSize,
	}, nil
}

func (r *Repository) GetAvatarIdByLink(userId uuid.UUID, link string) *uuid.UUID {
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
		log.Debugf("Invalid fragments while parsing picture link %s", fragments)
		return nil
	}
	avatarId, err := uuid.Parse(fragments[3])
	if err != nil {
		log.Debugf("Invalid picture ID while parsing picture link %s", link)
		return nil
	}
	return &avatarId
}
