package entity

import "github.com/google/uuid"

type UserInfo struct {
	UserId      uuid.UUID
	FirstName   *string
	LastName    *string
	Description *string
	AvatarUrl   *string
}
