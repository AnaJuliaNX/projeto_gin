package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Criação do token de acesso, a parte mais importante
type JWTService interface {
	GenerateToken(nome string, admin bool) string         //método para a geração de token
	ValidateToken(tokenString string) (*jwt.Token, error) //método para a validação do token
}

// Declarações personalizadas
type jwtCustomClaims struct {
	Nome  string `json:"nome"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

// Função para criar uma chave secreta
func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "projetogin.com",
	}
}

// Função para buscar a chave secreta criada
func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

// Função para gerar um token e usando uma bibliteca pronta
// Podem ser encontrados dados sobre ela no site https://jwt.io/libraries?language=Go
func (jwtSrv *jwtService) GenerateToken(username string, admin bool) string {
	claims := &jwtCustomClaims{
		username, //recebo o nome do usuário
		admin,    //recebo se ele é admin ou não
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(), //expiração do token
			Issuer:    jwtSrv.issuer,                         //emissor
			IssuedAt:  time.Now().Unix(),                     //hora em que o token foi criado
		},
	}
	//Cria um token com as informações
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//Assinando o token com a chave secreta
	t, erro := token.SignedString([]byte(jwtSrv.secretKey))
	if erro != nil {
		panic(erro)
	}
	return t
}

// Função de validação do token criado,vai analisar e validar ele
func (jwtSrv *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//validação do método de assinatura
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Método inesperado: %v", token.Header["alg"])
		}
		//Retorna a chave de assinatura secreta
		return []byte(jwtSrv.secretKey), nil
	})
}
