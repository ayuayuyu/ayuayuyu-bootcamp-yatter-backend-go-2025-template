package timelines

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	ui_errors "yatter-backend-go/app/ui/api/pkg/errors"
	"yatter-backend-go/app/usecase/timelines"
	"yatter-backend-go/pkg/errors"
)

// テストしやすいように、ハンドラーのインターフェースを定義
type Handler interface {
	GetTimelines(w http.ResponseWriter, r *http.Request)
}

func NewTimelinesHandler(timelinesUseCase timelines.GetPublicByTimelinesUseCase) Handler {
	return &timelinesHandlerImpl{
		timelinesUseCase: timelinesUseCase,
	}
}

var _ Handler = (*timelinesHandlerImpl)(nil)

// timelinesHandlerはtimelines関連のAPIを管理
type timelinesHandlerImpl struct {
	timelinesUseCase timelines.GetPublicByTimelinesUseCase
}

func (tl *timelinesHandlerImpl) GetTimelines(
	w http.ResponseWriter, r *http.Request) {

	onlyImageStr := r.URL.Query().Get("only_image")
	onlyImage, err := strconv.ParseBool(onlyImageStr)
	if err != nil {
		onlyImage = false
	}

	offsetStr := r.URL.Query().Get("offset")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}
	fmt.Printf("offsetStr: %v\n", offsetStr)
	fmt.Printf("offset: %v\n", offset)

	limitStr := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 0
	}
	fmt.Printf("limitStr: %v\n", limitStr)
	fmt.Printf("limit: %v\n", limit)

	ctx := r.Context()

	AllYweets, err := tl.timelinesUseCase.GetTimelines(ctx, onlyImage, offset, limit)
	if err != nil {
		ui_errors.Handle(w, err)
		return
	}

	// レスポンスに変換
	resp := toGetTimelinesResponse(AllYweets)

	// レスポンスをエンコードして返す
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		ui_errors.Handle(w, errors.ErrInternal.WithDevMessage(fmt.Sprintf("failed to encode response: %s", err.Error())))
		return
	}
}
