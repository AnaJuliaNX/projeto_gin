package controller

import (
	"net/http"
	"projeto_gin/service"
	"projeto_gin/tipos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Delega uma fubção ou nesse caso várias funções
type VideoControlller interface {
	FindAll() []tipos.Video
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) VideoControlller {
	validate = validator.New()
	return &controller{
		service: service,
	}
}

// Função para encontrar todos os videos cadastrados
func (c *controller) FindAll() []tipos.Video {
	return c.service.FindAll()
}

// Função para salvar os videos
func (c *controller) Save(ctx *gin.Context) error {
	var video tipos.Video
	erro := ctx.ShouldBindJSON(&video)
	//lido com o erros caso ocorra algum quando estiver salvadndo o video
	if erro != nil {
		return erro
	}

	erro = validate.Struct(video)
	if erro != nil {
		return erro
	}
	c.service.Save(video)
	return nil
}

func (c *controller) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()
	data := gin.H{
		"titulo": "Video Page",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
