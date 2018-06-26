package model
// Correlation model
// id1(post_id)  -   id2(tag_id)
type Correlation struct {
	Model
	ID1				uint64			`json:"id1"`
	ID2				uint64			`json:"id2"`
}