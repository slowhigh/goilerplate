package dto

type CreateMemo struct {
	Content string `json:"content" binding:"required"`
}
