package middlewares

import (
	"log"
	"net/http"
	"projeto_gin/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Validação da autorização do token da requisição http, returna um erro 401 se não for valido
func Autorizacao() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Autorização")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		//método de validação do token
		token, erro := service.NewJWTService().ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[Nome]: ", claims["nome"])
			log.Println("Claims[Admin]: ", claims["admin"])
			log.Println("Claims[Issuer]: ", claims["iss"])
			log.Println("Claims[IssuedAt]: ", claims["iat"])
			log.Println("Claims[ExpiresAt]: ", claims["exp"])
		} else {
			//Se não for válido exibo o erro
			log.Println(erro)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
