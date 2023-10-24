package tipos

import "time"

// Toda a validação é feita por aqui mesmo
type Pessoa struct {
	ID int64 `gorm:"primay_key;auto_increment" json:"id"`
	//campo obrigatório
	PrimeiroNome string `json:"primeironome" binding:"required"  gorm:"type:varchar(30)"`
	//campo obrigatório
	UltimoNome string `json:"ultimonome" binding:"required"  gorm:"type:varchar(30)"`
	//entre 10 a 130 anos
	Idade int64 `json:"idade" binding:"gte=18,lte=130"`
	//é validado, obrigatório e espera um email completo
	Email string `json:"email" binding:"required,email"  gorm:"type:varchar(256)"`
}

type Video struct {
	ID int64 `gorm:"primay_key;auto_increment" json:"id"`
	//minimo de 2 e máximo de 30 caracteres.
	//Só é um titulo válido se for preenchido com o que está no validator
	Titulo string `json:"titulo" binding:"min=2,max=40"  gorm:"type:varchar(40)"`
	//sem minimo e máximo de 40 caracteres
	Descricao string `json:"descricao" binding:"max=40"  gorm:"type:varchar(100)"`
	//campo obrigatório e espera uma url completa
	URL string `json:"url" binding:"required,url"  gorm:"type:varchar(256);UNIQUE"`
	//campo é obrigatório
	Autor    Pessoa `json:"autor" binding:"required"  gorm:"foreignkey:PessoaID"`
	PessoaID int64  `json:"-"`
	//Armazena a hora toda vez que um video é criado
	CreatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	//Armazena a hora toda vez que um livro é atualizado
	UpdatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
