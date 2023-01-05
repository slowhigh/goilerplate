package dto

type UpdateMemo struct {
	ID      string `json:"id" binding:"required"`
	Content string `json:"content" binding:"required"`
}
