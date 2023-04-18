package entity

type User struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Document  string `json:"document" gorm:"uniqueIndex;not null"`
	Email     string `json:"email" gorm:"uniqueIndex;not null"`
	Name      string `json:"name" gorm:"uniqueIndex"`
	Birthdate string `json:"birthdate"`
}