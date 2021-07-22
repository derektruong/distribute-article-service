package news

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func HeadLinesHandler(cl *Client) fiber.Handler {
	return func(c *fiber.Ctx) error{
		page := c.Query("page")

		if page == "" {
			page = "1"
		}

		results, err := cl.FetchThings("others", "", page,"https://newsapi.org/v2/top-headlines?country=us&page=%s&pageSize=%d&apiKey=%s")
		// fmt.Println(results)

		if err != nil {
			return c.Status(http.StatusInternalServerError).Render("general/notfound", nil)
		}

		currentPage, err := strconv.Atoi(page)

		if err != nil {
			return c.Status(http.StatusInternalServerError).Render("test.tmpl", nil)
		}

		resArray := make([][]Article, 0)

		crr := 0

		for {
			mini := int(math.Min(float64(len(results.Articles)), float64(crr + 3)))
			resArray = append(resArray, results.Articles[crr: mini])
			if mini == len(results.Articles) {
				break
			}
			crr += 3
		}

		search := &Search{
			Type: "Top Headlines",
			Path: "headlines",
			Query: "",
			CurrentPage: currentPage,
			TotalPages: int(math.Ceil(float64(results.TotalResults) / float64(cl.PageSize))),
			Results: results,
			RowResults: resArray,
		}

		return c.Status(http.StatusOK).Render("news/category", search)
	}
	
}
