package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/home", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})
	router.Run(":8080")
}
