package entity

import "github.com/google/uuid"

type UserMinimalInfo struct {
	UserId     uuid.UUID
	FullName   *string
	AvatarId   *uuid.UUID
	AvatarLink *string
}

type UserInfo struct {
	UserId      uuid.UUID
	FirstName   *string
	LastName    *string
	Description *string
	AvatarId    *uuid.UUID
	AvatarLink  *string
}
