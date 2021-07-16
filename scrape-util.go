package main

import (
	"encoding/json"
	"fmt"

	// "io/ioutil"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)


type article struct {
	Title 		string	`json:"title"`
	URL 		string	`json:"URL"`
	Score 		string 	`json:"score"`
	ScoreNum	string	`json:"scoreNum"`
	Comments 	string 	`json:"comments"`
	CommentsNum string	`json:"commentsNum"`
	Poster 		string	`json:"poster"`
	
}

type ranking struct {
	CurrentRank 		int	`json:"currentRank"`
	URL 				string	`json:"URL"`
}

func main() {
	var articles []article
	var pageCount int
	i := 0
	
	// Instantiate default collector
	c := colly.NewCollector()
	cssSelector := "tbody > tr:nth-child(3) > td > table > tbody"
	


	c.OnHTML(cssSelector, func(e *colly.HTMLElement) {
				e.ForEach("tr", func(_ int, h *colly.HTMLElement) {
					var art article
					// var rank ranking
					title := h.ChildText("td.title > a") 
					score := h.ChildText("td.subtext > span.score")
					
					scoreNum := strings.TrimRight(score, " points")
					scoreInt, err := strconv.Atoi(scoreNum)
					if err == nil {
							fmt.Println("Score Int: ")
							fmt.Println(scoreInt)
					}

					comments := h.ChildText("td.subtext > a:last-child")
					fmt.Println(comments)
					commentsNum := strings.TrimRight(comments, " comments")
					commentsNum = strings.Trim(commentsNum, " ")
					fmt.Println("Comments Num: ")
					fmt.Println(commentsNum)
					commentsInt, err := strconv.Atoi(commentsNum)
					// fmt.Println("Comments Int: ")
					// temp := strconv.Itoa(int(commentsInt))
					// fmt.Println(temp)
					// fmt.Println(strconv.Itoa(int(420)))
					if err == nil {
							fmt.Println("Comments Int: ")
							fmt.Println(strconv.Itoa(commentsInt))
					}
					
					currentRank := (scoreInt + 1) * (commentsInt + 1)
					fmt.Println("Current Rank: ")
					fmt.Println(currentRank)

					
					// currentRank := 
					// fmt.Println("One")
					// fmt.Println(articles[0])
					if title == "More" {
						c.Visit("news.ycombinator.com/" + h.ChildAttr("td.title > a", "href"))
						// fmt.Println("Two")
						// fmt.Println(articles[0])
					} else if score != "" {
						articles[i].Score = score
						articles[i].ScoreNum = scoreNum
						articles[i].Comments = comments
						articles[i].CommentsNum = commentsNum
						articles[i].Poster = h.ChildText("td.subtext > a.hnuser")
						
						// fmt.Println("Three")
						// fmt.Println(articles[0])
						i++
					} else if title != "" {
						art.Title = title
						art.URL = h.ChildAttr("td.title > a", "href")
						articles = append(articles, art)

						// rank.CurrentRank
						// fmt.Println("Four")
						// fmt.Println(articles[0])
					}
					
				})

	})
	// Before making a request print "Visiting ..."
	c.OnResponse(func(r *colly.Response) {
		pageCount++
        urlVisited := r.Request.URL
        fmt.Println(fmt.Sprintf("%d  DONE Visiting : %s", pageCount, urlVisited))
	})

	// Start scraping on https://news.ycombinator.com
	c.Visit("https://news.ycombinator.com/")
	fmt.Println(articles[0])
	fmt.Println("We did it!")

	// Marshal instances of articles and conert to JSON
	articleJSON, _ := json.Marshal(articles)
	// fmt.Println(string(articleJSON))
	articleJSONString := string(articleJSON)
	// fmt.Println(articleJSONString)
	writeJSONToFile(articleJSONString)
}


func writeJSONToFile(articleJSONString string) error {
	// fmt.Println("Five")
	// fmt.Println(articleJSONString)
    file, err := os.Create("output.json")
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = io.WriteString(file, articleJSONString)
    if err != nil {
        return err
    }
    return file.Sync()

	
}
func scoreJSONToFile(scoreJSONString string) error {

	file, err := os.Create("score.json")
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.WriteString(file, scoreJSONString)
		if err != nil {
			return err
		}
		return file.Sync()

}
// func comparePointsComments(articleJSONString string) error {
// 	forEach(articles) {

// 	}
// }