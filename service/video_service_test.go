package service

import (
	"projeto_gin/tipos"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Exemplo de informações para o teste
const (
	TITULO    = "Tudo sobre livros"
	DESCRICAO = "Video falanso sobre livros"
	URL       = "http://yotube.com/tudosobrelivros"
)

// Buscando os dados do video
func getVideos() tipos.Video {
	return tipos.Video{
		Titulo:    TITULO,
		Descricao: DESCRICAO,
		URL:       URL,
	}
}

// fazendo o teste do video
func TestFindAll(t *testing.T) {
	service := New()

	service.Save(getVideos()) //Buscando os salvos

	videos := service.FindAll()

	primeiroVideo := videos[0]
	assert.NotNil(t, videos)                            //os dados não podem ser nulos
	assert.Equal(t, TITULO, primeiroVideo.Titulo)       //O titulo do teste precisa ser no modelo salvo
	assert.Equal(t, DESCRICAO, primeiroVideo.Descricao) //A descrição precisa ser no modelo salvo
	assert.Equal(t, URL, primeiroVideo.URL)             //A url precisa ser no modelo salvo
}
