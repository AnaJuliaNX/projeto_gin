package main

import (
	"io"
	"net/http"
	"os"
	"projeto_gin/controller"
	"projeto_gin/middlewares"
	"projeto_gin/service"

	"github.com/gin-gonic/gin"
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

	//todas as rotas da minha API
	apiRoutes := server.Group("/api", middlewares.AutenticacaoBasic())
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

	viewRotas := server.Group("/view")
	{
		viewRotas.GET("/videos", VideoControlller.ShowAll)
	}

	//Inicialização do servidor com o Dockerfile
	port := os.Getenv("PORTA")
	if port == "" {
		port = "5000"
	}
	server.Run(":" + port)

	/*rotas do HTML e CSS
	Se quiser ver isso na web é só abrir uma aba do navegador e digitar a rota
	Nesse caso a rota é: http://localhost:8080/views/videos
	Lembrar de verificar se ta no servidor correto

	server.Run(":8080")
	*/

}
