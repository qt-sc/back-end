package model

type User struct {
 
	Id int64 `json:"id,omitempty" gorm:"id"`

	Name string `json:"name,omitempty" gorm:"name"`

	Password string `json:"password,omitempty" gorm:"password"`
 
	Articles []Article `json:"articles,omitempty" gorm:"articles"`
 
	Email string `json:"email,omitempty" gorm:"email"`

	Url string `json:"url,omitempty" gorm:"url"`
}