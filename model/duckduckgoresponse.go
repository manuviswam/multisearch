package model

type DuckDuckGoTopic struct {
	FirstURL string
	Text     string
}

type DuckDuckGoResponse struct {
	RelatedTopics []DuckDuckGoTopic
}
