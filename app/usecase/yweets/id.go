package yweets

import (
	"context"
	"yatter-backend-go/app/domain/object/yweets"
	"yatter-backend-go/app/usecase/query"
)

type GetYweetsByIdUseCase interface {
	GetYweetsById(ctx context.Context, id uint64) (*yweets.Yweets, error)
}

type YweetsQuerySeviceImpl struct {
	yweetsRepo query.YweetsId
}

func NewYweetsByIdUseCase(
	yweetsRepo query.YweetsId,
) *YweetsQuerySeviceImpl {
	return &YweetsQuerySeviceImpl{
		yweetsRepo: yweetsRepo,
	}
}

func (y *YweetsQuerySeviceImpl) GetYweetsById(
	ctx context.Context, id uint64) (*yweets.Yweets, error) {

	yweet, err := y.yweetsRepo.FindById(ctx, id)

	if err != nil {
		return nil, err
	}

	return yweet, nil
}
