package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/someday-94/TypeGoMongo-Server/controller"
	"github.com/someday-94/TypeGoMongo-Server/dto"
)

type VideoApi struct {
	loginController controller.LoginController
	videoController controller.VideoController
}

func NewVideoAPI(loginController controller.LoginController, videoController controller.VideoController) *VideoApi {
	return &VideoApi{
		loginController: loginController,
		videoController: videoController,
	}
}

// Path Information

// Authenticate godoc
// @Summary Provides a JSON Web Token
// @Description Authenticates a user and provides a JWT to Authorize API calls
// @ID Authentication
// @Consume application/x-www-form-urlencoded
// @Produce json
// @Param username formData string true "User credentials"
// @Param password formData string true "User credentials"
// @Success 200 {object} dto.JWT
// @Failure 401 {object} dto.Response
// @Router /auth/token [post]
func (api *VideoApi) Authenticate(ctx *gin.Context) {
	token := api.loginController.Login(ctx)
	if token != "" {
		ctx.JSON(http.StatusOK, &dto.JWT{
			Token: token,
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, &dto.Response{
			Message: "Not Authorized",
		})
	}
}

// GetVideo godoc
// @Security bearerAuth
// @Summary List existing videos
// @Description Get all the existing videos
// @Tags videos,list
// @Accept json
// @Produce json
// @Success 200 {array} entity.Video
// @Failure 401 {object} dto.Response
// @Router /videos [get]
func (api *VideoApi) GetVideos(ctx *gin.Context) {
	ctx.JSON(200, api.videoController.FindAll())
}

// CreateVideo godoc
// @Security bearerAuth
// @Summary Create new videos
// @Description Create a new video
// @Tags videos,create
// @Accept json
// @Produce json
// @Param video body entity.Video true "Create video"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /videos [post]
func (api *VideoApi) CreateVideo(ctx *gin.Context) {
	err := api.videoController.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &dto.Response{
			Message: "Success!",
		})
	}
}

// UpdateVideo godoc
// @Security bearerAuth
// @Summary Update videos
// @Description Update a single video
// @Tags videos
// @Accept json
// @Produce json
// @Param  id path int true "Video ID"
// @Param video body entity.Video true "Update video"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /videos/{id} [put]
func (api *VideoApi) UpdateVideo(ctx *gin.Context) {
	err := api.videoController.Update(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &dto.Response{
			Message: "Success!",
		})
	}
}

// DeleteVideo godoc
// @Security bearerAuth
// @Summary Remove videos
// @Description Delete a single video
// @Tags videos
// @Accept json
// @Produce json
// @Param  id path int true "Video ID"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /videos/{id} [delete]
func (api *VideoApi) DeleteVideo(ctx *gin.Context) {
	err := api.videoController.Delete(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &dto.Response{
			Message: "Success!",
		})
	}
}
