package service

import "github.com/someday-94/TypeGoMongo-Server/entity"

type VideoService interface {
	FindAll() []entity.Video
	Save(entity.Video) entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) FindAll() []entity.Video {
	return service.videos
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}
