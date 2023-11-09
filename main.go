package main

import (
	"fmt"
	"io"
	"net/http"
	"restapi/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Use(middleware.Authenticate, middleware.ResponseMiddleware)

	auth := gin.BasicAuth(gin.Accounts{
		"user": "pass",
	})
	adminGroup := router.Group("/admin", auth)
	adminGroup.GET("/getRequest", getData)

	clientGroup := router.Group("/client")
	clientGroup.GET("/getRequestQuery", getRequestQuery)

	router.POST("/postRequest", postRequest)

	router.GET("/getRequestParam/:name/:age", getRequestParam)
	//router.Run()

	server := &http.Server{
		Addr:         ":8081",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server.ListenAndServe()

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
