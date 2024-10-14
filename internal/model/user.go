package model

type User struct {
	Id       uint   `gorm:"primary_key"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
	Username string `gorm:"column:username"`
}

func (User) TableName() string {
	return "users"
}
