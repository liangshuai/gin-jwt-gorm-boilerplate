package model

const (
	CorrelationUserRole = iota
	CorrelationPostTag
)
// Correlation model
// id1(post_id)  -   id2(tag_id)
// id1(user_id)  -   id2(role_id)
type Correlation struct {
	Model
	ID1				uint16			`json:"id1"`
	ID2				uint16			`json:"id2"`

	Type			int				`json:"type"`
}