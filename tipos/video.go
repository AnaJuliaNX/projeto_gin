package tipos

// Toda a validação é feita por aqui mesmo
type Pessoa struct {
	//campo obrigatório
	PrimeiroNome string `json:"primeironome" binding:"required"`
	//campo obrigatório
	UltimoNome string `json:"ultimonome" binding:"required"`
	//entre 10 a 130 anos
	Idade int64 `json:"idade" binding:"gte=18,lte=130"`
	//é validado, obrigatório e espera um email completo
	Email string `json:"email" binding:"required,email"`
}

type Video struct {
	//minimo de 2 e máximo de 30 caracteres.
	//Só é um titulo válido se for preenchido com o que está no validator
	Titulo string `json:"titulo" binding:"min=2,max=40" validate:"terra-do-nunca"`
	//sem minimo e máximo de 40 caracteres
	Descricao string `json:"descricao" binding:"max=40"`
	//campo obrigatório e espera uma url completa
	URL string `json:"url" binding:"required,url"`
	//campo é obrigatório
	Autor Pessoa `json:"autor" binding:"required"`
}
