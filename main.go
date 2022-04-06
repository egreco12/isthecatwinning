package main

import (
	"fmt"
        "github.com/gocolly/colly/v2"
	"strconv"
)


type Player struct {
    Name string
    TotalScore int
    RoundOneScore int
    RoundTwoScore int
    RoundThreeScore int
    RoundFourScore int
}

func main() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML(".Table__TBODY", func(e *colly.HTMLElement) {
		fp := &Player{}
		p := &Player{}
		e.ForEachWithBreak("tr", func(idx int, row *colly.HTMLElement) bool {
			if (idx == 0) {
				fp = parseFirstPlace(row)
			}

			p = parseRowForPlayer("Tiger Woods", row)
			if (p.Name != "") {
				return false
			}

			return true
		})
		fmt.Println("first place:", fp)
		fmt.Println("input player:", p)
		fmt.Printf("Player %s beat player %s by %d strokes\n", fp.Name, p.Name, -1*(fp.TotalScore - p.TotalScore))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.espn.com/golf/leaderboard?tournamentId=2241")
}

func parseFirstPlace(e *colly.HTMLElement) *Player {
	newPlayer := Player{}
	done := false
	e.ForEachWithBreak("td", func(idx int, column *colly.HTMLElement) bool {
		if (idx == 2) {
			newPlayer.Name = column.Text
		}

		// idx 3 is total score
		if (idx == 3) {
			newPlayer.TotalScore, _ = strconv.Atoi(column.Text)
			done = true
		}
		if (done) {
			return false
		}

		return true
	})

	return &newPlayer
}

// Hackily go row by row to find the player instead of figuring out a query for it and return a Player
func parseRowForPlayer(player string, e *colly.HTMLElement) *Player {
	newPlayer := Player{}
	found := false
	done := false
	e.ForEachWithBreak("td", func(idx int, column *colly.HTMLElement) bool {
		if (idx == 2) && (column.Text == player) {
			newPlayer.Name = column.Text
			found = true
		}

		// idx 3 is total score
		if ((idx == 3) && (found)) {
			newPlayer.TotalScore, _ = strconv.Atoi(column.Text)
			done = true
		}
		if (done) {
			return false
		}

		return true
	})

	return &newPlayer
}
