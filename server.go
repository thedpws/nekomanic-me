package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	
	// Serve static files
	r.Static("/static", "./static")

	// Serve the HTML page
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/api/posts", func (c *gin.Context) {
		posts := []string{"Post 1", "Post 2", "Post 3"}
		c.HTML(http.StatusOK, "posts.html", gin.H{"posts": posts})
	})
	r.Run(":8080")
}
