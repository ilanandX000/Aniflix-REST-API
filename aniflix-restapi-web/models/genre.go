package models

type Genre struct {
	ID        int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string `gorm:"type:varchar(255)" json:"name"`
	CreatedAt int    `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int    `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedBy int    `json:"created_by"`
	UpdatedBy int    `json:"updated_by"`
}

func (Genre) TableName() string {
	return "genre"
}