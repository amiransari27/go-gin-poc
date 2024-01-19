package service

import "go-gin-api/src/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func NewVideo() VideoService {
	return &videoService{
		videos: make([]entity.Video, 0),
	}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}
func (service *videoService) FindAll() []entity.Video {
	return service.videos
}
