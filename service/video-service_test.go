package service

import (
	"testing"

	"github.com/someday-94/TypeGoMongo-Server/entity"
	"github.com/someday-94/TypeGoMongo-Server/repository"
	"github.com/stretchr/testify/assert"
)

const (
	TITLE       = "Video Title"
	DESCRIPTION = "Video Description"
	URL         = "https://www.youtube.com/"
)

func getVideo() entity.Video {
	return entity.Video{
		Title:       TITLE,
		Description: DESCRIPTION,
		URL:         URL,
	}
}

func TestFindAll(t *testing.T) {
	videoRepository := repository.NewVideoRepository()
	defer videoRepository.CloseDB()

	service := New(videoRepository)

	service.Save(getVideo())

	videos := service.FindAll()

	firstVideo := videos[0]
	assert.NotNil(t, firstVideo)
	assert.Equal(t, TITLE, firstVideo.Title)
	assert.Equal(t, DESCRIPTION, firstVideo.Description)
	assert.Equal(t, URL, firstVideo.URL)
}
