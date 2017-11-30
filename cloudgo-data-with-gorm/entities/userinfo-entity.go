package entities

import (
	"time"
)

// UserInfo .
type UserInfo struct {
	UID        int        `gorm:"column:uid;primary_key;AUTO_INCREMENT"`
	Username   string     `gorm:";type:varchar(100);not null;unique"`
	Departname string     `gorm:";type:varchar(100)"`
	CreatedAt  *time.Time `gorm:";not null"`
}

// NewUserInfo .
func NewUserInfo(u UserInfo) *UserInfo {
	if len(u.Username) == 0 {
		panic("Username shold not null!")
	}
	if u.CreatedAt == nil {
		t := time.Now()
		u.CreatedAt = &t
	}
	return &u
}
