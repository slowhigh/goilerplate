package services

import (
	"github.com/oxyrinchus/goilerplate/lib"
	"github.com/oxyrinchus/goilerplate/models"
	"github.com/oxyrinchus/goilerplate/repositories"
)

type MemoService struct {
	logger         lib.Logger
	memoRepository repositories.MemoRepository
}

// NewMemoService initialize memo service.
func NewMemoService(logger lib.Logger, memoRepository repositories.MemoRepository) MemoService {
	return MemoService{
		logger:         logger,
		memoRepository: memoRepository,
	}
}

// CreateMemo inserts the memo.
func (ms MemoService) CreateMemo(memo models.Memo) error {
	return ms.memoRepository.Create(memo)
}

// GetMemoByID gets the memo matching given the id.
func (ms MemoService) GetMemoByID(id string) (models.Memo, error) {
	return ms.memoRepository.FindOne(models.Memo{ID: id})
}

// GetMemosByUserID get all memos matching given the user's ID.
func (ms MemoService) GetMemosByUserID(userID string) (memos []models.Memo, err error) {
	return ms.memoRepository.FindAll(models.Memo{UserID: userID})
}

// UpdateMemo updates the memo matching the given memo. but the memo must contain an ID.
func (ms MemoService) UpdateMemo(memo models.Memo) error {
	return ms.memoRepository.Update(memo)
}

// DeleteUser deletes the memo matching the given memo's ID.
func (ms MemoService) DeleteMemo(memo models.Memo) error {
	return ms.memoRepository.Delete(memo)
}
