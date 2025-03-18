package api

type API struct {
	UserID       int
	AccessToken  string
	RefreshToken string
}

func NewAPI(userID int, accessToken string, refreshToken string) *API {
	return &API{UserID: userID, AccessToken: accessToken, RefreshToken: refreshToken}
}
