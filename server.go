/*
	This file launches a Gin web server for serving OTTOMH files.

	To change the port the server is hosted on, set the $PORT environment var.
*/

package main

import (
	"os"

	"github.com/cis3296f22/ottomh/backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/cis3296f22/ottomh/backend/types"
	"sync"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()

	r.Static("/static", "build/static")
	r.StaticFile("/favicon.ico", "build/favicon.ico")
	r.StaticFile("/robots.txt", "build/robots.txt")
	r.StaticFile("/manifest.json", "build/manifest.json")
	r.StaticFile("/logo192.png", "build/logo192.png")
	r.StaticFile("/logo512.png", "build/logo512.png")

	r.LoadHTMLFiles("build/index.html")

	r.GET("/echo", routes.EchoHandler)
	r.GET("/", routes.IndexHandler)
	// Catch-all route to work nicely with react-router
	r.GET("/:path", routes.IndexHandler)

	lob := types.World{Mu:sync.Mutex{}}
	r.POST("/CreateLobby", lob.CreateLobby)

	r.Run(":" + port)
}
