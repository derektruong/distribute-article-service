package handler

import (
	"errors"
	"net/http"

	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gofiber/fiber/v2"
)

type Result struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Count   int    `json:"count"`
	Quotes  []Quote `json:"quotes"`
}

type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
	Tag    string `json:"tag"`
}

func FetchQuotes(c *http.Client) (*Result, error) {
	req, _ := http.NewRequest("GET", "https://goquotes-api.herokuapp.com/api/v1/random?count=5", nil)

	req.Header.Add("Accept", "application/json")

	resp, err := c.Do(req)

	if err != nil {
		return nil, errors.New("errored when sending request to the server")
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}

	var res Result
	json.Unmarshal(body, &res)
	return &res, json.Unmarshal(body, &res)
}

func (q *Quote) FormatText() string {
	if len(q.Text) > 150 {
		description := q.Text[0:150]

		return fmt.Sprint(description + "...")
	}
	return fmt.Sprint(q.Text)
}

func QuotesHandler(cl *http.Client) fiber.Handler {
	return func(c *fiber.Ctx) error{
		results, err := FetchQuotes(cl)
		

		if err != nil {
			return c.Status(http.StatusInternalServerError).Render("general/notfound", nil)
		}


		return c.Status(http.StatusOK).Render("index", fiber.Map{
			"Quo1": results.Quotes[0:2],
			"TextCard1": results.Quotes[2].FormatText(),
			"AuthorCard1": results.Quotes[2].Author,
			"TextCard2": results.Quotes[3].FormatText(),
			"AuthorCard2": results.Quotes[3].Author,
			"TextCard3": results.Quotes[4].FormatText(),
			"AuthorCard3": results.Quotes[4].Author,
			// "Title": "Hello",
		})
	}
	
}