package service

// Pequena autenticação usando apenas o nome e senha
type LoginService interface {
	Login(username string, password string) bool
}

type loginService struct {
	authorizedUsername string
	authorizedPassword string
}

// Aqui normalmente seria feita uma consulta no banco de dados para verificação da veracidade dos dados
func NewLoginService() LoginService {
	return &loginService{
		authorizedUsername: "projeto-gin",
		authorizedPassword: "essaeasenha",
	}
}

func (service *loginService) Login(username string, password string) bool {
	return service.authorizedUsername == username &&
		service.authorizedPassword == password
}
