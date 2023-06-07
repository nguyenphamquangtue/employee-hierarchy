package dto

type User struct {
	ID       int    `gorm:"primaryKey"`
	Username string `gorm:"size:255"`
	Password string `gorm:"size:255"`
}
