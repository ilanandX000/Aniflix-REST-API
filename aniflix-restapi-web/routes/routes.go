package routes

import (
	"aniflix-restapi-web/controllers"

	"github.com/gin-gonic/gin"
)

func AniflixRoutes(router *gin.Engine) {
    animeGroup := router.Group("/animes")
    {
        animeGroup.GET("/", controllers.GetAllAnimes)
        animeGroup.GET("/:id", controllers.GetAnimeByID)
        animeGroup.POST("/", controllers.CreateAnime)
        animeGroup.PUT("/:id", controllers.UpdateAnime)
        animeGroup.DELETE("/:id", controllers.DeleteAnime)
    }

	AnimeEpisodeGroup := router.Group("/anime-episodes")
	{
		AnimeEpisodeGroup.GET("/", controllers.GetAllEpisodes)
		AnimeEpisodeGroup.GET("/:id", controllers.GetEpisodeByID)
		AnimeEpisodeGroup.POST("/", controllers.CreateEpisode)
		AnimeEpisodeGroup.PUT("/:id", controllers.UpdateEpisode)
		AnimeEpisodeGroup.DELETE("/:id", controllers.DeleteEpisode)
	}

	GenreGroup := router.Group("/genres")
	{
		GenreGroup.GET("/", controllers.GetAllGenres)
		GenreGroup.GET("/:id", controllers.GetGenreByID)
		GenreGroup.POST("/", controllers.CreateGenre)
		GenreGroup.PUT("/:id", controllers.UpdateGenre)
		GenreGroup.DELETE("/:id", controllers.DeleteGenre)
	}

	AnimeGenreGroup := router.Group("/anime-genres")
	{
		AnimeGenreGroup.GET("/", controllers.GetAllAnimeGenres)
		AnimeGenreGroup.GET("/:id", controllers.GetAnimeGenreByID)
		AnimeGenreGroup.POST("/", controllers.CreateAnimeGenre)
		AnimeGenreGroup.PUT("/:id", controllers.UpdateAnimeGenre)
		AnimeGenreGroup.DELETE("/:id", controllers.DeleteAnimeGenre)
	}


}
