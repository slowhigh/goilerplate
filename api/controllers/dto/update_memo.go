package dto

type UpdateMemo struct {
	Content string `json:"content" binding:"required"`
}
