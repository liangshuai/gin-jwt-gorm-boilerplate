package model

type Tag struct {
	Model
	Title				string			`gorm:"size:128" json:"title"`
	PostCount		int				`json:"postCount"`
}