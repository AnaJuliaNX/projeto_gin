package middlewares

import "github.com/gin-gonic/gin"

func AutenticacaoBasic() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"projeto_gin": "essaeasenha", //credencias: usuario e senha
	})
}
