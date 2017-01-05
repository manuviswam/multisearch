package model

type GoogleResponseItem struct {
	Link    string
	Snippet string
}

type GoogleResponse struct {
	Items []GoogleResponseItem
}
