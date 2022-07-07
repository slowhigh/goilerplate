package repository

import (
	"github.com/someday-94/TypeGoMongo-Server/entity"
)

type VideoRepository interface {
	Save(video entity.Video)
	Update(video entity.Video)
	Delete(video entity.Video)
	FindAll() []entity.Video
}

type videoRepository struct {
	db Database
}

func NewVideoRepository(db Database) VideoRepository {
	db.AutoMigrate(&entity.Video{}, &entity.Person{})
	
	return &videoRepository{
		db: db,
	}
}

func (videoRepo *videoRepository) Save(video entity.Video) {
	videoRepo.db.conn.Create(&video)
}

func (videoRepo *videoRepository) Update(video entity.Video) {
	videoRepo.db.conn.Save(&video)
}

func (videoRepo *videoRepository) Delete(video entity.Video) {
	videoRepo.db.conn.Delete(&video)
}

func (videoRepo *videoRepository) FindAll() []entity.Video {
	var videos []entity.Video
	videoRepo.db.conn.Set("gorm:auto_preload", true).Find(&videos)
	return videos
}
