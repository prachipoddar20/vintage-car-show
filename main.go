package main

import (
	"fmt"
	"vincs/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//version 1
	v1 := r.Group("v1")
	{
		///<summary>
		/// Computes minimum disparity for the speed given for all cars for all the test cases
		///</summary>
		///<param>models.input</param>
		///<returns>Returns map[string]int64 with status code 200 when request is successful</returns>
		v1.POST("/minDisparity", controllers.FindMinimumDisparityV1)
	}
	//version 2
	v2 := r.Group("v2")
	{
		///<summary>
		/// Fetches input from the url passed as query parameter and computes minimum disparity
		///</summary>
		///<param>url</param>
		///<returns>Returns map[string]int64 with status code 200 when request is successful</returns>
		v2.GET("/minDisparity", controllers.FindMinimumDisparityV2)
	}

	r.Run() //http://localhost:8080
	fmt.Println("Server is running")
}
