package controller

import (
	"projeto_gin/dto"
	"projeto_gin/service"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jWtService   service.JWTService
}

func NewLoginController(loginService service.LoginService, jWtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credenciais dto.Credenciais
	erro := ctx.ShouldBind(&credenciais)
	if erro != nil {
		return ""
	}
	//Chama o loginService com as credenciais de acesso
	autenticado := controller.loginService.Login(credenciais.Username, credenciais.Password)
	if autenticado {
		//Se estiver autenticado chama o cerviço de geração de token
		return controller.jWtService.GenerateToken(credenciais.Username, true)
	}
	return ""
}
