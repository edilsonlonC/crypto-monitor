package userModel


type User struct {
	Id  uint `gorm:"primaryKey"`
	Name string
	Surname string
	Username string
	Password string
}
