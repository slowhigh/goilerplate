package model

type User struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name"`
}

type UserInput struct {
	Name string `json:"name"`
}
