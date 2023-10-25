package repositorio

import (
	"projeto_gin/tipos"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type VideoRepositorio interface {
	Save(video tipos.Video)
	Update(video tipos.Video)
	Delete(video tipos.Video)
	FindAll() []tipos.Video
	CloseDB()
}

type bancodedados struct {
	connection *gorm.DB
}

func NewVideoRepositorio() VideoRepositorio {
	db, erro := gorm.Open("mysql", "youtuber:senhadele@/videos?charset=utf8&parseTime=True&loc=Local")
	if erro != nil {
		panic("Falha ao fazer a conexão com o banco de dados")
	}
	db.AutoMigrate(&tipos.Video{}, &tipos.Pessoa{})
	return &bancodedados{
		connection: db,
	}
}

// Função para fcehar a conexão com o banco de dados
func (db *bancodedados) CloseDB() {
	erro := db.connection.Close()
	if erro != nil {
		panic("Erro ao fechar o banco de dados")
	}
}

func (db *bancodedados) Save(video tipos.Video) {
	db.connection.Create(&video)
}

func (db *bancodedados) Update(video tipos.Video) {
	db.connection.Save(&video)
}

func (db *bancodedados) Delete(video tipos.Video) {
	db.connection.Delete(&video)
}

func (db *bancodedados) FindAll() []tipos.Video {
	var videos []tipos.Video
	//Dessa forma eu retorno tanto os dados dos videos como o da pessoa que ta vinculada ao video
	db.connection.Set("gorm:auto_preload", true).Find(&videos)
	return videos
}
