package gork

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"
	"github.com/gocolly/colly/v2/queue"
	"golang.org/x/time/rate"
)

// Search returns a list of search results from Google.
func Search(ctx context.Context, searchTerm string, opts ...SearchOptions) ([]Result, error) {

    if ctx == nil {
		ctx = context.Background()
	}

    rateLimit := rate.NewLimiter(rate.Inf, 0)
	if err := rateLimit.Wait(ctx); err != nil {
		return nil, err
	}

	c := colly.NewCollector(colly.MaxDepth(1))
	if len(opts) == 0 {
		opts = append(opts, SearchOptions{})
	}

    c.UserAgent = opts[0].UserAgent

	q, _ := queue.New(
		2,
		&queue.InMemoryQueueStorage{MaxSize: 10000},
	)

	results := []Result{}
	nextPageLink := ""

	var rErr error
	rank := 1

	c.OnRequest(func(r *colly.Request) {
		if err := ctx.Err(); err != nil {
			r.Abort()
			rErr = err
			return
		}

		if opts[0].FollowLinks == true && nextPageLink != "" {
			req, err := r.New("GET", nextPageLink, nil)
			if err == nil {
				q.AddRequest(req)
			}
		}
	})

    /* on error, set global error object */
	c.OnError(func(r *colly.Response, err error) {
		rErr = err
	})

    /*
        On div:
        look for potential link & save it in Result[]
    */
	c.OnHTML("div.g", func(e *colly.HTMLElement) {

		sel := e.DOM

		linkHref, _ := sel.Find("a").Attr("href")
		linkText := strings.TrimSpace(linkHref)
		titleText := strings.TrimSpace(sel.Find("div > div > div > a > h3").Text())

		if linkText != "" && linkText != "#" && titleText != "" {
			result := Result{
				URL:         linkText,
				Title:       titleText,
			}
			results = append(results, result)
			rank += 1
		}

		/*
            check if there is a next button at the end.
            if so, add it to, let's try that at the end of the current page parsing
        */
		nextPageHref, _ := sel.Find("a #pnnext").Attr("href")
		nextPageLink = strings.TrimSpace(nextPageHref)
	})

	c.OnHTML("div.g", func(e *colly.HTMLElement) {

		sel := e.DOM

		/* check if there is a next button at the end. */
		nextPageHref, exists := sel.Attr("href")
		if exists == true {
			start := getUrlStart(strings.TrimSpace(nextPageHref))
			nextPageLink = buildUrl(searchTerm, start)
			q.AddURL(nextPageLink)
		} else {
			nextPageLink = ""
		}
	})

	url := buildUrl(searchTerm, 0)

    /* set proxy for next search query */
	if opts[0].ProxyAddr != "" {
		proxySwitcher, err := proxy.RoundRobinProxySwitcher(opts[0].ProxyAddr)
		if err != nil {
			return nil, err
		}
		c.SetProxyFunc(proxySwitcher)
	}

	q.AddURL(url)
	q.Run(c)

    return results, rErr
}

func getUrlStart(uri string) int {
	u, err := url.Parse(uri)
	if err != nil {
		fmt.Println(err)
	}
	q := u.Query()
	ss := q.Get("start")
	si, _ := strconv.Atoi(ss)
	return si

}

func buildUrl(term string, start int) string {
    const baseURL = "https://www.google.com/search?q="

    term = strings.Trim(term, " ")
	term = strings.Replace(term, " ", "+", -1)
    if start != 0 {
        return fmt.Sprintf("%s%s&num%d", baseURL, term, start)
    }
    return fmt.Sprintf("%s%s", baseURL, term)
}
