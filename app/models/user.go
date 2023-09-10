package models

type User struct {
	ID
	Username string `json:"username" gorm:"not null;index;comment:用户名"`
	Password string `json:"password" gorm:"not null;default:'';comment:密码"`
	Timestamps
	SoftDelete
}
