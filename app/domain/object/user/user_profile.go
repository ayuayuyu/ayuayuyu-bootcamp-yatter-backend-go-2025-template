package user

import "time"

type UserProfile struct {
	ID             uint64
	Username       string
	DisplayName    string
	PasswordHash   string
	CreatedAt      time.Time
	FollowersCount uint64
	FollowingCount uint64
	Note           string
	Avatar         string
	Header         string
}
