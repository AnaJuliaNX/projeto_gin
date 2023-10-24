package controller

import (
	"net/http"
	"projeto_gin/service"
	"projeto_gin/tipos"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Delega uma fubção ou nesse caso várias funções
type VideoControlller interface {
	FindAll() []tipos.Video
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
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

// Funcão para atualizar um video, buscado pelo ID na rota
func (c *controller) Update(ctx *gin.Context) error {
	var video tipos.Video
	erro := ctx.ShouldBind(&video)
	if erro != nil {
		return erro
	}

	id, erro := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if erro != nil {
		return erro
	}
	video.ID = id

	erro = validate.Struct(video)
	if erro != nil {
		return erro
	}
	c.service.Update(video)
	return nil
}

// função para deleter um video, buscando pelo ID na rota
func (c *controller) Delete(ctx *gin.Context) error {
	var video tipos.Video
	id, erro := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if erro != nil {
		return erro
	}
	video.ID = id
	c.service.Delete(video)
	return nil
}

// Função para mostrar todos os videos que foram previamente salvos
func (c *controller) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()
	data := gin.H{
		"titulo": "Video Page",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
