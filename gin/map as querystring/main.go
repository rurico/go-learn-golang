package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// /post?ids[a]=1234&ids[b]=hello
//names[first]=thinkerou
//names[second]=tianou

func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		fmt.Printf("ids: %v; names: %v", ids, names)
	})
	router.Run(":8080")
}
