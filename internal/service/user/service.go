package user

import (
	"github.com/mephistolie/chefbook-backend-user/internal/service/dependencies/repository"
)

type Service struct {
	repo repository.User
	s3   repository.S3
}

func NewService(
	repo repository.User,
	s3 repository.S3,
) *Service {
	return &Service{
		repo: repo,
		s3:   s3,
	}
}
