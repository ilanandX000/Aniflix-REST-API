package models

type AnimeEpisode struct {
	ID        int    `gorm:"primaryKey;autoIncrement" json:"id"`
	AnimeID   int    `gorm:"not null" json:"anime_id"`
	Name      string `json:"name"`
	Sequence  int    `json:"sequence"`
	StreamURL string `json:"stream_url"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
	CreatedBy int    `json:"created_by"`
	UpdatedBy int    `json:"updated_by"`
}

func (AnimeEpisode) TableName() string {
	return "anime_episode"
}
