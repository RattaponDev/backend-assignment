package main

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)
func countBeef(text string) map[string]int {
	re := regexp.MustCompile(`[.,]`)
	cleanText := re.ReplaceAllString(text, "")
	words := strings.Fields(cleanText)
	beefCount := make(map[string]int)
	for _, word := range words {
		word = strings.ToLower(word)
		beefCount[word]++
	}
	return beefCount
}

func fetchData() (string, error) {
	response, err := http.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	r := gin.Default()
	r.GET("/beef/summary", func(c *gin.Context) {
		text, err := fetchData()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch data",
			})
			return
		}

		beefSummary := countBeef(text)

		c.JSON(http.StatusOK, gin.H{
			"beef": beefSummary,
		})
	})

	r.Run(":80")
}
