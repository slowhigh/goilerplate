package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	//"github.com/go-playground/validator/v10"
	"github.com/someday-94/TypeGoMongo-Server/entity"
	"github.com/someday-94/TypeGoMongo-Server/service"
	//"github.com/someday-94/TypeGoMongo-Server/validators"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

//var validate *validator.Validate

func New(service service.VideoService) VideoController {
	//validate = validator.New()
	//validate.RegisterValidation("is-cool", validators.ValidateCoolTitle) // "is-cool"이라는 이름의 유효성 검사 메소드를 추가

	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	//err = validate.Struct(video)
	//if err != nil {
	//	return err
	//}

	c.service.Save(video)
	return nil
}

func (c *controller) Update(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0) 
	if err != nil {
		return err
	}
	video.ID = id

	//err = validate.Struct(video)
	//if err != nil {
	//	return err
	//}

	c.service.Update(video)
	return nil
}

func (c *controller) Delete(ctx *gin.Context) error {
	var video entity.Video
	
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0) 
	if err != nil {
		return err
	}
	video.ID = id

	c.service.Delete(video)
	return nil
}

func (c *controller) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()
	data := gin.H{
		"title" : "Video Page",
		"videos" : videos,
	}

	ctx.HTML(http.StatusOK, "index.html", data)
}
