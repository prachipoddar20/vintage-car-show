package controllers

import (
	"net/http"
	"vincs/models"
	"vincs/service"

	"github.com/gin-gonic/gin"
)

func FindMinimumDisparityV1(c *gin.Context) {
	//declare variables
	var input models.Input
	var result = make(map[string]int64)

	//take input from request body
	if err := c.BindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}

	//call service to compute answer
	result = service.FindMinDisparity(input)

	//return result
	c.JSON(http.StatusOK, result)
}
