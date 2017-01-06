package model

type TwitterTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type Status struct {
	Text string
}

type TwitterResponse struct {
	Statuses []Status
}
