package configinfrastructurerepositorymodel

import "time"

type Config struct {
	ConfigID  int64     `gorm:"primaryKey;not null"`
	Name      string    `gorm:"column:name;type:varchar;not null"`
	Value     string    `gorm:"column:value;type:varchar;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;not null"`
	DeletedAt *time.Time
	Status    bool    `gorm:"default:false;not null"`
	Notes     *string `gorm:"column:notes;type:text"`
}

func (Config) TableName() string {
	return "config"
}
