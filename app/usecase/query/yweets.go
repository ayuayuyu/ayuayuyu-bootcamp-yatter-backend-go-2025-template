package query

import (
	"context"
	"yatter-backend-go/app/domain/object/yweets"
)

type YweetsId interface {
	FindById(ctx context.Context, id uint64) (*yweets.Yweets, error)
}
