package reader

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestName1(t *testing.T) {

	object, err := NewObject(`{
		"@context": "https://www.w3.org/ns/activitystreams",
		"type": "Note",
		"name": "A simple note"
		}`)

	t.Log(spew.Sdump(object))

	assert.Nil(t, err)
	assert.Equal(t, "Note", object.Type())
	assert.Equal(t, "A simple note", object.Name(""))
	assert.Equal(t, "A simple note", object.Name("en"))
	assert.Equal(t, "A simple note", object.Name("es"))
}

func TestName2(t *testing.T) {

	object, err := NewObject(`{
		"@context": "https://www.w3.org/ns/activitystreams",
		"type": "Note",
		"nameMap": {
			"en": "A simple note",
			"es": "Una nota sencilla",
			"zh-Hans": "一段简单的笔记"
		}
		}`)

	assert.Nil(t, err)
	assert.Equal(t, object.Type(), "Note")
	assert.Equal(t, object.Name(""), "A simple note")
	assert.Equal(t, object.Name("en"), "A simple note")
	assert.Equal(t, object.Name("es"), "Una nota sencilla")
	assert.Equal(t, object.Name("zh-Hans"), "一段简单的笔记")

}
