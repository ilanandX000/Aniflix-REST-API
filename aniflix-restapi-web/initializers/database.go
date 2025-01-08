package initializers

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    var err error
    dsn := "digienz:digienz888@tcp(35.237.219.245)/aniflix?parseTime=True"
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to the database")
    }
    //DB.AutoMigrate(&models.Anime{}, &models.AnimeEpisode{}, &models.Genre{}, &models.AnimeGenre{})
}
