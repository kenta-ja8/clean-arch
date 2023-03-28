package entity

import "time"

type User struct {
	Id       string    `validate:"required,len=36"`
	Name     string    `validate:"required,max=255"`
	Birthday time.Time `validate:"required,lte"`
}

func (u *User) Age() int {
	today := time.Now()
	age := today.Year() - u.Birthday.Year()
	if today.YearDay() < u.Birthday.YearDay() {
		age--
	}
	return age
}

type Book struct {
	Id    string `validate:"required,len=36"`
	Title string `validate:"required,max=255"`
}
