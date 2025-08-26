package timelines

import (
	"context"
	"yatter-backend-go/app/domain/object/yweets"
	"yatter-backend-go/app/usecase/query"
)

type GetPublicByTimelinesUseCase interface {
	GetTimelines(
		ctx context.Context, onlyImage bool, offset int, limit int) ([]*yweets.Yweets, error)
}

type TimelinesQuerySeviceImpl struct {
	timelinesRepo query.TimelinesLimit
}

func NewTimelinesUseCase(
	timelinesRepo query.TimelinesLimit,
) *TimelinesQuerySeviceImpl {
	return &TimelinesQuerySeviceImpl{
		timelinesRepo: timelinesRepo,
	}
}
func (tl *TimelinesQuerySeviceImpl) GetTimelines(
	ctx context.Context, onlyImage bool, offset int, limit int) ([]*yweets.Yweets, error) {

	timelines, err := tl.timelinesRepo.AllYweets(ctx, onlyImage, offset, limit)
	if err != nil {
		return nil, err
	}
	return timelines, nil
}
