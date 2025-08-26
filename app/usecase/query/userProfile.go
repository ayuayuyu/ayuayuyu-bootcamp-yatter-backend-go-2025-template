package query

import (
	"context"
	"yatter-backend-go/app/domain/object/user"
)

type UserProfile interface {
	FindByUsername(ctx context.Context, username string) (*user.UserProfile, error)
}
