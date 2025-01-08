package models

type AnimeGenre struct {
	ID        int `gorm:"primaryKey;autoIncrement" json:"id"`
	AnimeID   int `gorm:"not null" json:"anime_id"`
	GenreID   int `gorm:"not null" json:"genre_id"`
	CreatedAt int `json:"created_at"`
	UpdatedAt int `json:"updated_at"`
	CreatedBy int `json:"created_by"`
	UpdatedBy int `json:"updated_by"`
}

func (AnimeGenre) TableName() string {
	return "anime_genre"
}
