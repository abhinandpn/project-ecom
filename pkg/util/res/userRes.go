package res

import "time"

type UserResStruct struct {
	ID          uint      `json:"id" copier:"must"`
	FName       string    `json:"first_name" copier:"must"`
	LName       string    `json:"last_name" copier:"must"`
	Age         uint      `json:"age" copier:"must"`
	Email       string    `json:"email" copier:"must"`
	UserName    string    `json:"user_name" copire:"must"`
	Number      string    `json:"phone" copier:"must"`
	BlockStatus bool      `json:"block_status" copier:"must"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt   time.Time `json:"updated_at"`
}
type UserFindWithUserName struct {
	UserName string `json:"user_name" gorm:"not null"`
}
