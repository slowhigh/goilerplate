package model

type NewMemo struct {
	Content string `json:"content"`
	UserId  string `json:"userId"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Memo struct {
	ID      string `json:"id" bson:"_id"`
	Content string `json:"content"`
	Author  *User  `json:"author"`
}
