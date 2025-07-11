package webappdto

type CreateUserRq struct {
	TelegramID    int64  `json:"telegram_id"`
	ReferalUserID int64  `json:"referal"`
	Username      string `json:"username"`
	Language      string `json:"language"`
}
type CreateUserRp struct {
	UserID int64 `json:"user_id"`
}

type GetUserRq struct {
	UserID int64 `param:"user_id"`
}
type GetUserRp struct {
	TelegramID   int64  `json:"telegram_id"`
	Username     string `json:"username"`
	AvatarLink   string `json:"avatar_link"`
	Balance      string `json:"balance"`
	Rating       int64  `json:"rating"`
	FriendsCount int    `json:"friends_count"`
}
