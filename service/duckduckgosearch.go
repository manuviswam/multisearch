package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	m "github.com/manuviswam/multisearch/model"
)

const (
	duckDuckGoUrl = "http://api.duckduckgo.com/?format=json&q=%s"
)

type DuckDuckGoSearch struct {
}

func (d *DuckDuckGoSearch) Search(query string) m.SearchResult {
	resp, err := http.Get(fmt.Sprintf(duckDuckGoUrl, query))
	if err != nil {
		return m.SearchResult{
			Error: err.Error(),
		}
	}
	defer resp.Body.Close()

	duckDuckGoResponse := m.DuckDuckGoResponse{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&duckDuckGoResponse)
	if err != nil {
		return m.SearchResult{
			Error: err.Error(),
		}
	}

	if len(duckDuckGoResponse.RelatedTopics) < 1 {
		return m.SearchResult{
			Error: "No response obtained",
		}
	}
	return m.SearchResult{
		Url:  duckDuckGoResponse.RelatedTopics[0].FirstURL,
		Text: duckDuckGoResponse.RelatedTopics[0].Text,
	}

}
