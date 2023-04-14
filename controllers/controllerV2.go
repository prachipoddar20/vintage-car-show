package controllers

import (
	"encoding/json"
	"net/http"
	"vincs/models"
	"vincs/service"

	"github.com/gin-gonic/gin"
)

func FindMinimumDisparityV2(c *gin.Context) {
	//declare variables
	var input models.Input
	var result = make(map[string]int64)

	//fetch input from link
	url, _ := c.GetQuery("url")
	urlResponse, _ := http.Get(url)

	//decode url response into json object
	json.NewDecoder(urlResponse.Body).Decode(&input)

	//call service to compute answer
	result = service.FindMinDisparity(input)

	//return result
	c.JSON(http.StatusOK, result)
}
