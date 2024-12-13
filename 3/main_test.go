package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)
func mockFetchData() (string, error) {
	return `Fatback t-bone t-bone, pastrami  ..   t-bone.  pork, meatloaf jowl enim.  Bresaola t-bone.`, nil
}

func TestCountBeef(t *testing.T) {
	text := `Fatback t-bone t-bone, pastrami  ..   t-bone.  pork, meatloaf jowl enim.  Bresaola t-bone.`
	expectedBeefCount := map[string]int{
		"t-bone": 4,
        "fatback": 1,
        "pastrami": 1,
        "pork": 1,
        "meatloaf": 1,
        "jowl": 1,
        "enim": 1,
        "bresaola": 1,
	}
	result := countBeef(text)
	
	assert.Equal(t, expectedBeefCount, result, "The word count should match the expected value")
}


func TestBeefSummaryEndpoint(t *testing.T) {
	r := gin.Default()
	r.GET("/beef/summary", func(c *gin.Context) {
		text, err := mockFetchData()
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

	req, _ := http.NewRequest("GET", "/beef/summary", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	expectedResponse := `{"beef":{"t-bone":4,"fatback":1,"pastrami":1,"pork":1,"meatloaf":1,"jowl":1,"enim":1,"bresaola":1}}`
	assert.JSONEq(t, expectedResponse, w.Body.String())
}
