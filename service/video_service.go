package service

import (
	"projeto_gin/repositorio"
	"projeto_gin/tipos"
)

// Interface que vai criar e adicionar novos videos
type VideoService interface {
	Save(tipos.Video) tipos.Video
	Update(video tipos.Video)
	Delete(video tipos.Video)
	FindAll() []tipos.Video
}

// Retorna no final um slice of videos
type videoService struct {
	videoRepositorio repositorio.VideoRepositorio
}

func New(repo repositorio.VideoRepositorio) VideoService {
	return &videoService{
		videoRepositorio: repo,
	}
}

func (service *videoService) Save(video tipos.Video) tipos.Video {
	service.videoRepositorio.Save(video)
	return video
}

func (service *videoService) Update(video tipos.Video) {
	service.videoRepositorio.Update(video)
}

func (service *videoService) Delete(video tipos.Video) {
	service.videoRepositorio.Delete(video)
}

func (service *videoService) FindAll() []tipos.Video {
	return service.videoRepositorio.FindAll()
}
