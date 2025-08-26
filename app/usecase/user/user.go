package user

import (
	"context"
	"yatter-backend-go/app/domain/object/user"
	"yatter-backend-go/app/usecase/query"
)

type GetSingleUserByUsernameUseCase interface {
	GetSingleUserByUsername(ctx context.Context, username string) (userData *user.UserProfile, err error)
}

type UserQuerySeviceImpl struct {
	userRepo query.UserProfile
}

func NewUserByUsernameUseCase(
	userRepo query.UserProfile,
) *UserQuerySeviceImpl {
	return &UserQuerySeviceImpl{
		userRepo: userRepo,
	}
}

func (uc *UserQuerySeviceImpl) GetSingleUserByUsername(
	ctx context.Context, username string) (userData *user.UserProfile, err error) {

	// 引数の username で、SELECTクエリを叩く
	// 叩いた結果を返す
	user, err := uc.userRepo.FindByUsername(ctx, username)

	// もし、ユーザが存在しなければエラーを返す
	if err != nil {
		return nil, err
	}

	//問題が無ければ、レスポンスデータとエラーnilを返す
	return user, nil
}
