package controller

import (
	"github.com/gin-gonic/gin"
	"ictm.uz/video/entity"
	"ictm.uz/video/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

// FindAll implements VideoController


func New(service service.VideoService) controller {
	return controller{
		service: service,
	}
}


func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

// Save implements VideoController
func (c *controller) Save(ctx *gin.Context)entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)

	c.service.Save(video)
	return video
}