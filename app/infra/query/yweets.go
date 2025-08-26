package query

import (
	"context"
	"time"
	"yatter-backend-go/app/domain/object/user"
	"yatter-backend-go/app/domain/object/yweets"
	"yatter-backend-go/app/usecase/query"

	"github.com/jmoiron/sqlx"
)

var _ query.YweetsId = (*YweetsQueryServiceImpl)(nil)

type YweetsQueryServiceImpl struct {
	db *sqlx.DB
}

func NewYweetsIdRepository(db *sqlx.DB) *YweetsQueryServiceImpl {
	return &YweetsQueryServiceImpl{db: db}
}

type FindByIdDTO struct {
	ID               uint64 `db:"id"`
	User             FindByUsernameDTO
	Content          string    `db:"content"`
	CreatedAt        time.Time `db:"created_at"`
	ImageAttachments ImageAttachments
}

type ImageAttachments struct {
	ID          uint64 `db:"id"`
	Type        string `db:"type"`
	Url         string `db:"url"`
	Description string `db:"description"`
}

func (y *YweetsQueryServiceImpl) FindById(
	ctx context.Context,
	id uint64,
) (*yweets.Yweets, error) {
	var dbId FindByIdDTO
	err := y.db.GetContext(ctx, &dbId,
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
        WHERE 
            yweet.id = ?`, id,
	)

	if err != nil {
		return nil, err
	}

	yweet := yweets.Yweets{
		ID: dbId.ID,
		User: user.UserProfile{
			ID:             dbId.User.ID,
			Username:       dbId.User.Username,
			DisplayName:    dbId.User.DisplayName,
			CreatedAt:      dbId.User.CreatedAt,
			FollowersCount: 0,
			FollowingCount: 0,
			Note:           dbId.User.Note,
			Avatar:         dbId.User.Avatar,
			Header:         dbId.User.Header,
		},
		Content:   dbId.Content,
		CreatedAt: dbId.CreatedAt,
		ImageAttachments: yweets.ImageAttachments{
			ID:          123,
			Type:        "image",
			Url:         "string",
			Description: "string",
		},
	}

	return &yweet, nil
}
