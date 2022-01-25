package reader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestID(t *testing.T) {

	object, err := UnmarshalJSON(`{
		"@context": "https://www.w3.org/ns/activitystreams",
		"name": "Foo",
		"id": "http://example.org/foo"
	  }`)

	assert.Nil(t, err)
	assert.Equal(t, "http://example.org/foo", object.ID())
}

func TestType1(t *testing.T) {

	object, err := UnmarshalJSON(`{
		"@context": "https://www.w3.org/ns/activitystreams",
		"summary": "A foo",
		"type": "http://example.org/Foo"
	  }`)

	assert.Nil(t, err)
	assert.Equal(t, "http://example.org/Foo", object.Type())
	assert.Equal(t, []string{"http://example.org/Foo"}, object.TypeArray())
}

func TestType2(t *testing.T) {

	object, err := UnmarshalJSON(`{
		"@context": "https://www.w3.org/ns/activitystreams",
		"summary": "A foo",
		"type": ["first", "second", "third"]
	  }`)

	assert.Nil(t, err)
	assert.Equal(t, "first", object.Type())
	assert.Equal(t, []string{"first", "second", "third"}, object.TypeArray())
}

func TestActor1(t *testing.T) {

	object, err := UnmarshalJSON(`{
		"@context": "https://www.w3.org/ns/activitystreams",
		"summary": "Sally offered the Foo object",
		"type": "Offer",
		"actor": "http://sally.example.org",
		"object": "http://example.org/foo"
	  }`)

	assert.Nil(t, err)
	assert.Equal(t, "http://sally.example.org", object.Actor())

	actor := object.ActorObject()
	assert.Equal(t, "", actor.Type())
	assert.Equal(t, "http://sally.example.org", actor.ID())
	assert.Equal(t, actor.Summary(), "")
	assert.Equal(t, "", actor.Summary("es"))
}

func TestActor2(t *testing.T) {

	object, err := UnmarshalJSON(`{
		"@context": "https://www.w3.org/ns/activitystreams",
		"summary": "Sally offered the Foo object",
		"type": "Offer",
		"actor": {
		  "type": "Person",
		  "id": "http://sally.example.org",
		  "summary": "Sally"
		},
		"object": "http://example.org/foo"
	  }`)

	assert.Nil(t, err)
	assert.Equal(t, object.Actor(), "http://sally.example.org")

	actor := object.ActorObject()
	assert.Equal(t, actor.Type(), "Person")
	assert.Equal(t, actor.ID(), "http://sally.example.org")
	assert.Equal(t, actor.Summary(), "Sally")
	assert.Equal(t, actor.Summary("es"), "Sally")
}

func TestAttachment(t *testing.T) {

	object, err := UnmarshalJSON(`{
		"@context": "https://www.w3.org/ns/activitystreams",
		"type": "Note",
		"name": "Have you seen my cat?",
		"attachment": [
		  {
			"type": "Image",
			"content": "This is what he looks like.",
			"url": "http://example.org/cat.jpeg"
		  }
		]
	  }`)

	assert.Nil(t, err)
	assert.Equal(t, "Note", object.Type())
	assert.Equal(t, "Have you seen my cat?", object.Name())
	assert.Equal(t, "http://example.org/cat.jpeg", object.Attachment())

	attachment := object.AttachmentObject()
	assert.Equal(t, "Image", attachment.Type())
	assert.Equal(t, "This is what he looks like.", attachment.Content())
	assert.Equal(t, "http://example.org/cat.jpeg", attachment.URL())
}

func TestName1(t *testing.T) {

	object, err := UnmarshalJSON(`{
		"@context": "https://www.w3.org/ns/activitystreams",
		"type": "Note",
		"name": "A simple note"
		}`)

	assert.Nil(t, err)
	assert.Equal(t, "Note", object.Type())
	assert.Equal(t, "A simple note", object.Name())
	assert.Equal(t, "A simple note", object.Name("en"))
	assert.Equal(t, "A simple note", object.Name("es"))
}

func TestName2(t *testing.T) {

	object, err := UnmarshalJSON(`{
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
	assert.Equal(t, object.Name(), "A simple note")
	assert.Equal(t, object.Name("en"), "A simple note")
	assert.Equal(t, object.Name("es"), "Una nota sencilla")
	assert.Equal(t, object.Name("zh-Hans"), "一段简单的笔记")

}

/*
func TestDuration1(t *testing.T) {

	object, err := UnmarshalJSON(`{
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
*/
