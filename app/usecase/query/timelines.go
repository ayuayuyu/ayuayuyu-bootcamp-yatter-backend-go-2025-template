package query

import (
	"context"
	"yatter-backend-go/app/domain/object/yweets"
)

type TimelinesLimit interface {
	AllYweets(ctx context.Context, onlyImage bool, offset int, limit int) ([]*yweets.Yweets, error)
}
