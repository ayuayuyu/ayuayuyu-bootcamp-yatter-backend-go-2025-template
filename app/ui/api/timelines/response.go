package timelines

import (
	"yatter-backend-go/app/domain/object/yweets"
	"yatter-backend-go/app/ui/api/user"
	response_yweet "yatter-backend-go/app/ui/api/yweets"
)

func toGetTimelinesResponse(tl []*yweets.Yweets) *[]response_yweet.GetYweetResponse {
	var responses []response_yweet.GetYweetResponse

	for _, yweet := range tl {
		responses = append(responses, response_yweet.GetYweetResponse{
			ID: yweet.ID,
			User: user.PostUsersResponse{
				ID:             yweet.User.ID,
				Username:       yweet.User.Username,
				DisplayName:    yweet.User.DisplayName,
				CreatedAt:      yweet.User.CreatedAt.Format("2006-01-02T15:04:05.000Z"),
				FollowersCount: int(yweet.User.FollowersCount),
				FollowingCount: int(yweet.User.FollowingCount),
				Note:           yweet.User.Note,
				Avatar:         yweet.User.Avatar,
				Header:         yweet.User.Header,
			},
			Content:   yweet.Content,
			CreatedAt: yweet.CreatedAt.Format("2006-01-02T15:04:05.000Z"),
			ImageAttachments: response_yweet.GetImageAttachmentsResponse{
				ID:          yweet.ImageAttachments.ID,
				Type:        yweet.ImageAttachments.Type,
				Url:         yweet.ImageAttachments.Url,
				Description: yweet.ImageAttachments.Description,
			},
		})
	}

	return &responses
}
