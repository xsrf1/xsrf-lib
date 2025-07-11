package dto

import "time"

// Create user
type CreateUserRq struct {
	UserID        int64  `json:"ti"`
	ReferalUserID int64  `json:"ref"`
	Username      string `json:"un"`
	Language      string `json:"lg"`
}
type CreateUserRp struct {
	TelegramID int64  `json:"i"`
	Username   string `json:"u"`
	AvatarLink string `json:"a"`

	Balance      *int64 `json:"b"`
	EulaAccepted *bool  `json:"e"`
	Rating       int64  `json:"r"`

	CachedFriends *int       `json:"c"`
	LastLogin     *time.Time `json:"l"`
}

// Get user
type GetUserByTelegramIDRq struct {
	TelegramID int64 `json:"ti"`
}
type GetUserByIDRq struct {
	UserID int64 `json:"ui"`
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

type SetRatingRq struct {
	TelegramID int64 `json:"ti"`
	Rating     int64 `json:"r"`
}
type SetRatingRp struct{}

// Разделить пользователя на разные объекты и получать их по отдельности
