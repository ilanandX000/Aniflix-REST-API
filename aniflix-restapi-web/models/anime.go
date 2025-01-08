package models

import (
	"time"
)

type Anime struct {
    ID          int      `gorm:"primaryKey;autoIncrement" json:"id"`
    Title       string    `gorm:"type:longtext" json:"title"`
    Synopsis    string    `gorm:"type:longtext" json:"synopsis"`
    Thumbnail   string    `gorm:"type:longtext" json:"thumbnail"`
    Rating      float32   `gorm:"type:float" json:"rating"`
    Status      int       `gorm:"type:bigint" json:"status"`
    ReleaseDate time.Time `gorm:"type:date" json:"release_date"`
    CreatedAt   int 	  `json:"created_at"`
    UpdatedAt   int	      `json:"updated_at"`
    CreatedBy   int       `json:"created_by"`
    UpdatedBy   int       `json:"updated_by"`


}
func (Anime) TableName() string {
    return "anime"
}
