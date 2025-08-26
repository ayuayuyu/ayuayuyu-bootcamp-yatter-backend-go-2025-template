package yweets

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	ui_errors "yatter-backend-go/app/ui/api/pkg/errors"
	"yatter-backend-go/app/usecase/yweets"
	"yatter-backend-go/pkg/errors"

	"github.com/go-chi/chi/v5"
)

// テストしやすいように、ハンドラーのインターフェースを定義
type Handler interface {
	GetYweetsById(w http.ResponseWriter, r *http.Request)
}

func NewYweetsHandler(yweetsIdFindUseCase yweets.GetYweetsByIdUseCase) Handler {
	return &yweetsHandlerImpl{
		yweetsIdFindUseCase: yweetsIdFindUseCase,
	}
}

var _ Handler = (*yweetsHandlerImpl)(nil)

// yweetsHandlerはyweets関連のAPIを管理
type yweetsHandlerImpl struct {
	yweetsIdFindUseCase yweets.GetYweetsByIdUseCase
}

// yweet Idで参照
func (y *yweetsHandlerImpl) GetYweetsById(
	w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	convertedStrUint64, _ := strconv.ParseUint(id, 10, 64)

	ctx := r.Context()

	// ユーザー名取得のユースケースを呼び出し
	gotYweets, err := y.yweetsIdFindUseCase.GetYweetsById(ctx, convertedStrUint64)
	if err != nil {
		ui_errors.Handle(w, err)
		return
	}

	// レスポンスに変換
	resp := toGetYweetsIdResponse(gotYweets)

	// レスポンスをエンコードして返す
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		ui_errors.Handle(w, errors.ErrInternal.WithDevMessage(fmt.Sprintf("failed to encode response: %s", err.Error())))
		return
	}
}
