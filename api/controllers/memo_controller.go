package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/oxyrinchus/goilerplate/api/controllers/dto"
	"github.com/oxyrinchus/goilerplate/api/controllers/msg"
	"github.com/oxyrinchus/goilerplate/common"
	"github.com/oxyrinchus/goilerplate/lib"
	"github.com/oxyrinchus/goilerplate/models"
	"github.com/oxyrinchus/goilerplate/services"
)

type MemoController struct {
	logger      lib.Logger
	userService services.UserService
	memoService services.MemoService
}

// NewMemoController initialize memo controller.
func NewMemoController(logger lib.Logger, userService services.UserService, memoService services.MemoService) MemoController {
	return MemoController{
		logger:      logger,
		userService: userService,
		memoService: memoService,
	}
}

// CreateMemo creates the memo.
func (mc MemoController) CreateMemo(c *gin.Context) {
	status, data, msg := func() (int, *models.Memo, string) {
		userID, exists := c.Get(common.CURRENT_USER_ID)
		if !exists {
			mc.logger.Error(common.ERR_EMPTY_USER)
			return http.StatusInternalServerError, nil, msg.CONTACT_SERVER_ADMIN
		}

		var dto dto.CreateMemo
		if err := c.ShouldBind(&dto); err != nil {
			mc.logger.Error(err)
			return http.StatusBadRequest, nil, msg.BAD_REQUEST
		}

		newMemo := models.Memo{
			ID:      uuid.New().String(),
			Content: dto.Content,
			UserID:  userID.(string),
		}

		if err := mc.memoService.CreateMemo(newMemo); err != nil {
			mc.logger.Error(err)
			return http.StatusInternalServerError, nil, msg.CONTACT_SERVER_ADMIN
		}

		mc.logger.Info(fmt.Sprintf("Create memo. [ID:%s, UserID:%s]", newMemo.ID, newMemo.UserID))
		return http.StatusOK, &newMemo, ""
	}()

	c.JSON(status, gin.H{"data": data, "msg": msg})
}

// FindOneMemo finds the memo.
func (mc MemoController) FindOneMemo(c *gin.Context) {
	status, data, msg := func() (int, *models.Memo, string) {
		memoID := c.Param("id")
		if memoID == "" {
			mc.logger.Error(common.ERR_EMPTY_PARAM)
			return http.StatusBadRequest, nil, msg.BAD_REQUEST
		}

		foundMemo, err := mc.memoService.GetMemoByID(memoID)
		if err != nil {
			mc.logger.Error(err)
			return http.StatusInternalServerError, nil, msg.CONTACT_SERVER_ADMIN
		}

		mc.logger.Info("FindOneMemo")
		return http.StatusOK, &foundMemo, ""
	}()

	c.JSON(status, gin.H{"data": data, "msg": msg})
}

// FindAllMemo finds all memos.
func (mc MemoController) FindAllMemo(c *gin.Context) {
	status, data, msg := func() (int, *[]models.Memo, string) {
		userID, exists := c.Get(common.CURRENT_USER_ID)
		if !exists {
			mc.logger.Error(common.ERR_EMPTY_USER)
			return http.StatusInternalServerError, nil, msg.CONTACT_SERVER_ADMIN
		}

		foundMemos, err := mc.memoService.GetMemosByUserID(userID.(string))
		if err != nil {
			mc.logger.Error(err)
			return http.StatusInternalServerError, nil, msg.CONTACT_SERVER_ADMIN
		}

		return http.StatusOK, &foundMemos, ""
	}()

	c.JSON(status, gin.H{"data": data, "msg": msg})
}

// UpdateMemo updates the memo.
func (mc MemoController) UpdateMemo(c *gin.Context) {
	status, data, msg := func() (int, *models.Memo, string) {
		userID, exists := c.Get(common.CURRENT_USER_ID)
		if !exists {
			mc.logger.Error(common.ERR_EMPTY_USER)
			return http.StatusInternalServerError, nil, msg.CONTACT_SERVER_ADMIN
		}

		var dto dto.UpdateMemo
		if err := c.ShouldBind(&dto); err != nil {
			mc.logger.Error(err)
			return http.StatusBadRequest, nil, msg.BAD_REQUEST
		}

		newMemo := models.Memo{
			ID:      dto.ID,
			Content: dto.Content,
			UserID:  userID.(string),
		}

		if err := mc.memoService.UpdateMemo(newMemo); err != nil {
			mc.logger.Error(err)
			return http.StatusInternalServerError, nil, msg.CONTACT_SERVER_ADMIN
		}

		return http.StatusOK, &newMemo, ""
	}()

	c.JSON(status, gin.H{"data": data, "msg": msg})
}

// DeleteMemo deletes the memo.
func (mc MemoController) DeleteMemo(c *gin.Context) {
	status, data, msg := func() (int, *string, string) {
		memoID := c.Param("id")
		if memoID == "" {
			mc.logger.Error(common.ERR_EMPTY_PARAM)
			return http.StatusBadRequest, nil, msg.BAD_REQUEST
		}

		userID, exists := c.Get(common.CURRENT_USER_ID)
		if !exists {
			mc.logger.Error(common.ERR_EMPTY_USER)
			return http.StatusInternalServerError, nil, msg.CONTACT_SERVER_ADMIN
		}

		err := mc.memoService.DeleteMemo(models.Memo{ID: memoID, UserID: userID.(string)})
		if err != nil {
			mc.logger.Error(err)
			return http.StatusInternalServerError, nil, msg.CONTACT_SERVER_ADMIN
		}

		mc.logger.Info("delete memo. ", "memoID : ", memoID)
		return http.StatusOK, &memoID, ""
	}()

	c.JSON(status, gin.H{"data": data, "msg": msg})
}
