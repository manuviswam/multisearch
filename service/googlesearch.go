package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	m "github.com/manuviswam/multisearch/model"
)

const (
	googleUrl = "https://www.googleapis.com/customsearch/v1?key=%s&cx=017576662512468239146:omuauf_lfve&num=1&q=%s"
)

type GoogleSearch struct {
	APIKey string
}

func (g *GoogleSearch) Search(query string, c chan m.SearchResult) {
	defer close(c)
	resp, err := http.Get(fmt.Sprintf(googleUrl, g.APIKey, query))
	if err != nil {
		c <- m.SearchResult{
			Error: err.Error(),
		}
		return
	}
	defer resp.Body.Close()

	googleResponse := m.GoogleResponse{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&googleResponse)
	if err != nil {
		c <- m.SearchResult{
			Error: err.Error(),
		}
		return
	}

	if len(googleResponse.Items) < 1 {
		c <- m.SearchResult{
			Error: "No response obtained",
		}
		return
	}
	c <- m.SearchResult{
		Url:  googleResponse.Items[0].Link,
		Text: googleResponse.Items[0].Snippet,
	}
}
