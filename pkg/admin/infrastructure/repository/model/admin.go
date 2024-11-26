package admininfrastructurerepositorymodel

import "time"

type Admin struct {
	AdminID   int64     `gorm:"primaryKey;not null"`
	Name      string    `gorm:"column:name;not null"`
	Email     string    `gorm:"column:email;type:varchar(255);not null;unique"`
	Password  string    `gorm:"column:password;type:text;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;not null"`
	DeletedAt *time.Time
	LastLogin *time.Time `gorm:"type:timestamptz"`
}

func (Admin) TableName() string {
	return "admin"
}
