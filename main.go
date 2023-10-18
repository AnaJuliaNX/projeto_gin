package main

import (
	"io"
	"net/http"
	"os"
	"projeto_gin/controller"
	"projeto_gin/middlewares"
	"projeto_gin/service"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService     service.VideoService        = service.New()
	VideoControlller controller.VideoControlller = controller.New(videoService)
)

// Faz com que salve em um arquivo "gin.log" todas as vezes em que rodo uma rota
func SetupOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	/*
		 //Criada apenas para saber se estava funcionando tudo certo
			server := gin.Default()
		   	server.GET("/test", func(ctx *gin.Context) {
		   		ctx.JSON(200, gin.H{
		   			"message": "OK!",
		   		})
		   	})
	*/

	SetupOutput()

	server := gin.New()

	//para carregar o arquivo CSS
	server.Static("/css", "./templates/css")

	//Para carregar os arquivos HTML
	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), middlewares.Logger(),
		middlewares.AutenticacaoBasic(), gindump.Dump())

	//todas as rotas da minha API
	apiRoutes := server.Group("/api")
	{
		// Para buscar os videos já cadastrados
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, VideoControlller.FindAll())
		})

		// Para cadastrar um video novo
		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			//Lido com os erros caso tiver algum enquanto estiver salvando os dados do video
			erro := VideoControlller.Save(ctx)
			if erro != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video inserido é valido"})
			}

		})
	}

	//rotas do HTML e CSS
	viewRutes := server.Group("/view")
	{
		viewRutes.GET("/videos", VideoControlller.ShowAll)
	}

	server.Run(":8080")
}
