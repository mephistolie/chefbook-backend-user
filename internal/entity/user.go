package entity

import "github.com/google/uuid"

type UserMinimalInfo struct {
	UserId      uuid.UUID
	DisplayName *string
	AvatarId    *uuid.UUID
	AvatarLink  *string
}

type UserInfo struct {
	UserId      uuid.UUID
	DisplayName *string
	Description *string
	AvatarId    *uuid.UUID
	AvatarLink  *string
}
