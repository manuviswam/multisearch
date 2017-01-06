package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	m "github.com/manuviswam/multisearch/model"
)

const (
	Url = "https://www.googleapis.com/customsearch/v1?key=%s&cx=017576662512468239146:omuauf_lfve&num=1&q=%s"
)

type GoogleSearch struct {
	APIKey string
}

func (g *GoogleSearch) Search(query string) m.SearchResult {
	resp, err := http.Get(fmt.Sprintf(Url, g.APIKey, query))
	if err != nil {
		return m.SearchResult{
			Error: err.Error(),
		}
	}
	defer resp.Body.Close()

	googleResponse := m.GoogleResponse{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&googleResponse)
	if err != nil {
		return m.SearchResult{
			Error: err.Error(),
		}
	}

	if len(googleResponse.Items) < 1 {
		return m.SearchResult{
			Error: "No response obtained",
		}
	}
	return m.SearchResult{
		Url:  googleResponse.Items[0].Link,
		Text: googleResponse.Items[0].Snippet,
	}

}
