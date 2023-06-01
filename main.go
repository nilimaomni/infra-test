package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/headers", func(c *gin.Context) {
		headers := c.Request.Header

		for key, values := range headers {
			for _, value := range values {
				fmt.Printf("%s: %s\n", key, value)
			}
		}

		c.String(200, "Request headers printed!")
	})

	// Run the server on port 8080
	router.Run(":8080")
}
