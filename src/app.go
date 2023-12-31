// Inside your app.go file

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Configure Gin to look for templates in the "templates" directory
	router.LoadHTMLGlob("src/templates/*.html")

	// Define a route that renders an HTML page
	router.GET("/", func(c *gin.Context) {
		// Use the HTML function to render an HTML response
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Simple Page",
		})
	})

	// Run the web server on port 8080
	router.Run(":8080")
}
