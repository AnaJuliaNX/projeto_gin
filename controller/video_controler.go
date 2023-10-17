package controller

import (
	"projeto_gin/service"
	"projeto_gin/tipos"

	"github.com/gin-gonic/gin"
)

type VideoControlller interface {
	FindAll() []tipos.Video
	Save(ctx *gin.Context) tipos.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoControlller {
	return controller{
		service: service,
	}
}

func (c *controller) FindAll() []tipos.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) tipos.Video {
	var video tipos.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video

}
