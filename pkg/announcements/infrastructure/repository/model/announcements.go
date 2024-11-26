package announcementsinfrastructurerepositorymodel

import "time"

type Announcements struct {
	AnnouncementsID int64     `gorm:"primaryKey;not null"`
	Level           string    `gorm:"column:level;gorm:ENUM('INFO','WARNING','ALERT')"`
	Title           string    `gorm:"column:title"`
	Description     string    `gorm:"column:description"`
	IsPriority      bool      `gorm:"column:is_priority"`
	CreatedAt       time.Time `gorm:"autoCreateTime;not null"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime;not null"`
	DeletedAt       *time.Time
	Notes           *string `gorm:"column:notes;type:varchar"`
}

func (Announcements) TableName() string {
	return "announcements"
}
