package main

import (
	"io/ioutil"
	"fmt"
        "github.com/gocolly/colly/v2"
	"strconv"
	"encoding/json"
)


type Player struct {
	Name string `json:"name"`
        TotalScore int `json:"totalScore"`
	Cut bool `json:"cut"`
}

func main() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML(".Table__TBODY", func(e *colly.HTMLElement) {
		fp := &Player{}
		tw := &Player{}
		js := &Player{}

		// lolll this is so bad, too lazy to refactor tho
		e.ForEachWithBreak("tr", func(idx int, row *colly.HTMLElement) bool {
			if (idx == 0) {
				fp = parseFirstPlace(row)
			}

			tw = parseRowForPlayer("Tiger Woods", row)
			if (tw.Name != "") {
				return false
			}

			return true
		})
		e.ForEachWithBreak("tr", func(idx int, row *colly.HTMLElement) bool {
			if (idx == 0) {
				fp = parseFirstPlace(row)
			}

			js = parseRowForPlayer("Sandy Lyle", row)
			if (js.Name != "") {
				return false
			}

			return true
		})
		fmt.Println("first place:", fp)
		fmt.Println("tw:", tw)
		fmt.Println("jordo:", js)
		fmt.Printf("Player %s beat player %s by %d strokes\n", fp.Name, tw.Name, -1*(fp.TotalScore - tw.TotalScore))
                var l []Player
		l = append(l, *fp)
		l = append(l, *tw)
		l = append(l, *js)
		file, _ := json.MarshalIndent(l, "", " ")
		_ = ioutil.WriteFile("players.json", file, 0644)
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
			if (column.Text == "E") {
				newPlayer.TotalScore = 0
			}
			s, e := strconv.Atoi(column.Text)
			if e != nil {
				newPlayer.TotalScore = s
			}
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
		if ((idx == 2) && (column.Text == player)) {
			newPlayer.Name = column.Text
			found = true
		}

		// idx 3 is total score
		if ((idx == 3) && (found)) {
			newPlayer.Cut = column.Text == "CUT"
			if (column.Text == "E") {
				newPlayer.TotalScore = 0
			} else {
				newPlayer.TotalScore, _ = strconv.Atoi(column.Text)
			}
			done = true
		}
		if (done) {
			return false
		}

		return true
	})

	return &newPlayer
}
