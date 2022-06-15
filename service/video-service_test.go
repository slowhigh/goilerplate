package service

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/someday-94/TypeGoMongo-Server/entity"
)

const (
	TITLE = "Video Title"
	DESCRIPTION = "Video Description"
	URL = "https://www.youtube.com/"
)

func getVideo() entity.Video {
	return entity.Video{
		Title: TITLE,
		Description: DESCRIPTION,
		URL: URL,
	}
}

func TestFindAll(t *testing.T) {
	service := New()

	service.Save(getVideo())

	videos := service.FindAll()

	firstVideo := videos[0]
	assert.Nil(t, videos)
	assert.Equal(t, TITLE, firstVideo.Title)
	assert.Equal(t,DESCRIPTION, firstVideo.Description)
	assert.Equal(t, URL, firstVideo.URL)
}