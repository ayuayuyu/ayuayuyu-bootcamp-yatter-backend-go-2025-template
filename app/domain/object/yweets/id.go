package yweets

import (
	"time"
	"yatter-backend-go/app/domain/object/user"
)

type Yweets struct {
	ID               uint64
	User             user.UserProfile
	Content          string
	CreatedAt        time.Time
	ImageAttachments ImageAttachments
}

type ImageAttachments struct {
	ID          uint64
	Type        string
	Url         string
	Description string
}
