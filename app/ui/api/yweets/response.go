package yweets

import (
	"yatter-backend-go/app/domain/object/yweets"
	"yatter-backend-go/app/ui/api/user"
)

type GetYweetResponse struct {
	ID               uint64 `json:"id"`
	User             user.PostUsersResponse
	Content          string `json:"content"`
	CreatedAt        string `json:"created_at"`
	ImageAttachments GetImageAttachmentsResponse
}

type GetImageAttachmentsResponse struct {
	ID          uint64 `json:"id"`
	Type        string `json:"type"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

func toGetYweetsIdResponse(y *yweets.Yweets) *GetYweetResponse {
	return &GetYweetResponse{
		ID: y.ID,
		User: user.PostUsersResponse{
			ID:             y.User.ID,
			Username:       y.User.Username,
			DisplayName:    y.User.DisplayName,
			CreatedAt:      y.User.CreatedAt.Format("2006-01-02T15:04:05.000Z"),
			FollowersCount: int(y.User.FollowersCount),
			FollowingCount: int(y.User.FollowingCount),
			Note:           y.User.Note,
			Avatar:         y.User.Avatar,
			Header:         y.User.Header,
		},
		Content:   y.Content,
		CreatedAt: y.CreatedAt.Format("2006-01-02T15:04:05.000Z"),
		ImageAttachments: GetImageAttachmentsResponse{
			ID:          y.ImageAttachments.ID,
			Type:        y.ImageAttachments.Type,
			Url:         y.ImageAttachments.Url,
			Description: y.ImageAttachments.Description,
		},
	}
}
