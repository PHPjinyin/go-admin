package models

import "strconv"

type User struct {
	ID
	Username string `json:"username" gorm:"not null;index;comment:用户名"`
	Password string `json:"-" gorm:"not null;default:'';comment:密码"`
	Mobile   string `json:"mobile" json:"mobile" binding:"required"`
	Timestamps
	SoftDelete
}

func (user User) GetUid() string {
	return strconv.Itoa(int(user.ID.ID))
}
