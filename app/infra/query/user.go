package query

import (
	"context"
	"time"
	"yatter-backend-go/app/domain/object/user"
	"yatter-backend-go/app/usecase/query"

	"github.com/jmoiron/sqlx"
)

var _ query.UserProfile = (*UserQueryServiceImpl)(nil)

type UserQueryServiceImpl struct {
	db *sqlx.DB
}

func NewUserProfileRepository(db *sqlx.DB) *UserQueryServiceImpl {
	return &UserQueryServiceImpl{db: db}
}

type FindByUsernameDTO struct {
	ID          uint64    `db:"id"`
	Username    string    `db:"username"`
	DisplayName string    `db:"display_name"`
	CreatedAt   time.Time `db:"created_at"`
	Note        string    `db:"note"`
	Avatar      string    `db:"avatar"`
	Header      string    `db:"header"`
}

func (uc *UserQueryServiceImpl) FindByUsername(
	ctx context.Context,
	username string,
) (*user.UserProfile, error) {
	var dbUsername FindByUsernameDTO
	err := uc.db.GetContext(ctx, &dbUsername,
		`SELECT id,username,display_name,created_at,note,avatar,header FROM user WHERE username = ?`, username,
	)

	if err != nil {
		return nil, err
	}

	userProfile := user.UserProfile{
		ID:             dbUsername.ID,
		Username:       dbUsername.Username,
		DisplayName:    dbUsername.DisplayName,
		CreatedAt:      dbUsername.CreatedAt,
		FollowersCount: 0,
		FollowingCount: 0,
		Note:           dbUsername.Note,
		Avatar:         dbUsername.Avatar,
		Header:         dbUsername.Header,
	}

	return &userProfile, nil
}
