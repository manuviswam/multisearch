package model

type SearchResult struct {
	Error string
	Url   string
	Text  string
}

type SearchResults struct {
	Google     SearchResult
	Twitter    SearchResult
	Duckduckgo SearchResult
}

type SearchResponse struct {
	Query   string
	Results SearchResults
}
