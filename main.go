package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func routes() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Home Page",
		})
	})
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{
			"title": "User Login",
		})
	})

	fmt.Println("server started on :8080")
	// define port `r.Run(":8080")`
	log.Fatal(router.Run(":8080"))
}

func main() {
	routes()
}
