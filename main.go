package main

import (
	"fmt"
        "github.com/gocolly/colly/v2"
)
func main() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML(".PlayerRow__Overview", func(e *colly.HTMLElement) {
		fmt.Println("Found player", e.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.espn.com/golf/leaderboard?tournamentId=2241")
}
