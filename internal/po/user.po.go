package po

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"column:id;type:char(36);primaryKey" json:"id"`
	Username  string    `gorm:"column:username;type:varchar(50);unique;not null" json:"username"`
	Password  string    `gorm:"column:password;type:varchar(100);not null" json:"password"`
	IsActive  bool      `gorm:"column:is_active;type:tinyint(1);default:1" json:"is_active"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Roles     []Role         `gorm:"many2many:roles" json:"roles"`
}

func (u *User) TableName() string {
	return "users"
}
