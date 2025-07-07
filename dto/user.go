package dto

import "time"

// Get user or create it if its not exists
type GetUserRq struct {
	TelegramID int64 `json:"ti"`
}

type GetUserRp struct {
	TelegramID int64  `json:"i"`
	Username   string `json:"u"`
	AvatarLink string `json:"a"`

	Balance      *int64 `json:"b"`
	EulaAccepted *bool  `json:"e"`
	Rating       int64  `json:"r"`

	CachedFriends *int       `json:"c"`
	LastLogin     *time.Time `json:"l"`
}

type AcceptEulaRq struct {
	TelegramID int64 `json:"ti"`
}
type AcceptEulaRp struct{}

// Разделить пользователя на разные объекты и получать их по отдельности
