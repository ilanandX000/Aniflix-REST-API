package main

import (
	"aniflix-restapi-web/initializers"
	"aniflix-restapi-web/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectDatabase()
}

func main() {
	router := gin.Default()

	routes.AniflixRoutes(router)

	router.Run(":8080")
}