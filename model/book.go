package model

type NewMemo struct {
	Content string `json:"content"`
	UserId  string `json:"userId"`
}

type User struct {
	ID   string `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

type Memo struct {
	ID      string `json:"id" bson:"_id" gorm:"primary_key"`
	Content string `json:"content" gorm:"type:varchar(512)"`
	Author  *User  `json:"author" gorm:"foreignkey:UserID"`
	UserID  string `json:"-"`
}
