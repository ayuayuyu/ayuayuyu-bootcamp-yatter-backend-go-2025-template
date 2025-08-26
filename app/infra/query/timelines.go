package query

import (
	"context"
	"yatter-backend-go/app/domain/object/user"
	"yatter-backend-go/app/domain/object/yweets"
	"yatter-backend-go/app/usecase/query"

	"github.com/jmoiron/sqlx"
)

var _ query.TimelinesLimit = (*TimeLinesQueryServiceImpl)(nil)

type TimeLinesQueryServiceImpl struct {
	db *sqlx.DB
}

func NewTimelinesRepository(db *sqlx.DB) *TimeLinesQueryServiceImpl {
	return &TimeLinesQueryServiceImpl{db: db}
}

func (tl *TimeLinesQueryServiceImpl) AllYweets(
	ctx context.Context,
	onlyImage bool, offset int, limit int) ([]*yweets.Yweets, error) {

	//offsetがない時
	if offset < 0 {
		offset = 0
	}
	//limitがない時
	if limit == 0 {
		limit = 80
	}

	var dbAll []FindByIdDTO
	err := tl.db.SelectContext(ctx, &dbAll,
		`SELECT
        yweet.id,
        yweet.content,
        yweet.created_at,
        user.id AS "user.id",
        user.username AS "user.username",
        user.display_name AS "user.display_name",
        user.created_at AS "user.created_at",
        user.note AS "user.note",
        user.avatar AS "user.avatar",
        user.header AS "user.header"
        FROM
        yweet
        JOIN user ON yweet.user_id = user.id
		ORDER BY yweet.created_at DESC
		LIMIT ? OFFSET ?`, limit, offset)

	if err != nil {
		return nil, err
	}

	timelines := []*yweets.Yweets{}

	for _, dto := range dbAll {
		timelines = append(timelines, &yweets.Yweets{
			ID: dto.ID,
			User: user.UserProfile{
				ID:             dto.User.ID,
				Username:       dto.User.Username,
				DisplayName:    dto.User.DisplayName,
				CreatedAt:      dto.User.CreatedAt,
				FollowersCount: 0,
				FollowingCount: 0,
				Note:           dto.User.Note,
				Avatar:         dto.User.Avatar,
				Header:         dto.User.Header,
			},
			Content:   dto.Content,
			CreatedAt: dto.CreatedAt,
			ImageAttachments: yweets.ImageAttachments{
				ID:          123,
				Type:        "image",
				Url:         "string",
				Description: "string",
			},
		})
	}

	return timelines, nil
}
