package model

type Memo struct {
	ID      string `json:"id" bson:"_id"`
	Content string `json:"content"`
	Author  *User  `json:"author"`
}

type MemoInput struct {
	Content string     `json:"content"`
	Author  *UserInput `json:"author"`
}
