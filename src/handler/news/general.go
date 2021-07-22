package news

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	http 		*http.Client
	key 		string
	PageSize 	int
}

type Result struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles 	 []Article `json:"articles"`

}

type Article struct {
	Source struct {
		ID   interface{} `json:"id"`
		Name string      `json:"name"`
	} `json:"source"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	URLToImage  string    `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
	Content     string    `json:"content"`
}

type Search struct {
	Type string
	Path string
	Query      string
	CurrentPage   int
	TotalPages int
	Results *Result
	RowResults    [][]Article
}

func NewClient(httpClient *http.Client, key string, pageSize int) *Client{
	if pageSize > 100 {
		pageSize = 100
	}

	return &Client{httpClient, key, pageSize}
}

func (c *Client) FetchThings(location, query, page, path string) (*Result, error) {
	var endpoint string
	if(location == "search") {
		endpoint = fmt.Sprintf(path, url.QueryEscape(query), page, c.PageSize, c.key)
	} else {
		endpoint = fmt.Sprintf(path, page, c.PageSize, c.key)
	}
	
	resp, err := c.http.Get(endpoint)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}

	var res Result
	json.Unmarshal(body, &res)
	return &res, json.Unmarshal(body, &res)
}

func (a *Article) FormatPublishedDate() string {
	year, month, day := a.PublishedAt.Date()

	return fmt.Sprint(day, month, ", ", year)
}

func (a *Article) FormatDescription() string {
	if len(a.Description) > 150 {
		description := a.Description[0:150]

		return fmt.Sprint(description + "...")
	}
	return fmt.Sprint(a.Description)
}

func (s *Search) IsFirstPage() int {
	if s.CurrentPage == 1 {
		return 1
	}
	return s.CurrentPage - 1
}

func (s *Search) IsLastPage() int {
	if s.CurrentPage >= s.TotalPages {
		return s.CurrentPage
	}

	return s.CurrentPage + 1
}