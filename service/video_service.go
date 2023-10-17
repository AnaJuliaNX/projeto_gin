package service

import "projeto_gin/tipos"

// Interface que vai criar e adicionar novos videos
type VideoService interface {
	Save(tipos.Video) tipos.Video
	FindAll() []tipos.Video
}

type videoService struct {
	videos []tipos.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(video tipos.Video) tipos.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []tipos.Video {
	return service.videos
}
