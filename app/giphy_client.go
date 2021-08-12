package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type GiphyData struct {
	EmbedURL string `json:"embed_url"`
}

type GiphyResp struct {
	Data GiphyData `json:"data"`
}

type GiphyClient struct {
	URL    string
	APIKey string
}

func (g GiphyClient) FetchGIF(searchTerm string) (string, error) {
	reqURL, err := url.Parse(g.URL)
	if err != nil {
		return "", err
	}

	q := reqURL.Query()
	q.Set("api_key", g.APIKey)
	q.Set("tag", searchTerm)
	q.Set("rating", "PG13")
	reqURL.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var gifResp GiphyResp
	err = json.NewDecoder(resp.Body).Decode(&gifResp)
	if err != nil {
		return "", err
	}

	fmt.Println(gifResp.Data.EmbedURL)
	return gifResp.Data.EmbedURL, nil
}
