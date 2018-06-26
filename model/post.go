package model

type Post struct {
	Model
	AuthorID				uint64			`json:"authorID" structs:"authorID"`
	Title					string			`gorm:"size:128" json:"title"  structs:"title"`
	Content					string			`gorm:"type:text" json:"content" structs:"content"`
	Status					string			`sql:"index" json:"status" structs:"status"`
	Topped					bool 			`json:"topped" structs:"status"`
	ViewCount				int				`json:"viewCount" structs:"viewCount"`
	CommentCount			int				`json:"commentCount" structs:"commentCount"`
	Tags					string			`gorm:"type:text" json:"tags" structs:"tags"`
}


const (
	ArticleStatusDraft = iota
	ArticleStatusAuditing
	ArticleStatusOk
	ArticleStatusDeleted
)