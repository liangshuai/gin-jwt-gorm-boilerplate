package model

type Comment struct {
	Model
	PostID				uint64				`json:"postID"`
	AuthorID			uint64				`json:"authorID"`
	Content				string				`gorm:"type:text" json:"content"`
}