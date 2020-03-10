package reader

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDuration1(t *testing.T) {

	object, err := NewObject(`{
		"@context": "https://www.w3.org/ns/activitystreams",
		"type": "Video",
		"name": "Birds Flying",
		"url": "http://example.org/video.mkv",
		"duration": "PT2H"
	  }`)

	assert.Nil(t, err)
	assert.Equal(t, object.Type(), "Video")
	assert.Equal(t, object.Duration(), time.Duration(5*time.Second))
}
