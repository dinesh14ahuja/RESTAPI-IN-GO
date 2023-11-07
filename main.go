package main

import (
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/getRequest", getData)
	router.POST("/postRequest", postRequest)
	router.GET("/getRequestQuery", getRequestQuery)
	router.GET("/getRequestParam/:name/:age", getRequestParam)
	router.Run()

}

func getRequestParam(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"name": fmt.Sprintf("Request name is %s", name),
		"age":  fmt.Sprintf("Request age is %s", age),
	})
}

func getRequestQuery(c *gin.Context) {
	name := c.Query("name")
	c.JSON(200, gin.H{
		"data": fmt.Sprintf("Hello from %s", name),
	})
}

func postRequest(c *gin.Context) {
	body := c.Request.Body
	value, _ := io.ReadAll(body)
	c.JSON(200, gin.H{
		"data":     "Post Request data",
		"bodydata": string(value),
	})
}
func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hello from golang",
	})
}
