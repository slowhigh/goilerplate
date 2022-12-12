package dto

type UpdateMemo struct {
	ID string `json:"id" binding:"required"`
	CreateMemo
}
