package main

import (
	"fmt"

	globals "home-owner/globals"
	middleware "home-owner/middleware"
	routes "home-owner/routes"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.html")

	router.Use(sessions.Sessions("session", cookie.NewStore(globals.Secret)))

	public := router.Group("/")
	routes.PublicRoutes(public)

	private := router.Group("/")
	private.Use(middleware.Auth)
	routes.PrivateRoutes(private)

	fmt.Println("server started on :8080")
	// define port `r.Run(":8080")`
	router.Run(":8080")
}
