package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/golang-gin-poc/entity"
	"gitlab.com/pragmaticreviews/golang-gin-poc/service"
	"gitlab.com/pragmaticreviews/golang-gin-poc/validators"
	"gopkg.in/go-playground/validator.v9"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.BindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}
