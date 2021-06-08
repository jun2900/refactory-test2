package models

type User struct {
	Id    string `json:"id" gorm:"unique_index;not null"`
	Email string `json:"email" gorm:"size:45;unique_index;not null"`
}
