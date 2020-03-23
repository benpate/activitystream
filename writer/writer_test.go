package writer

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestWriter(t *testing.T) {

	object := Accept(
		Person("John Connor", "en").ID("john@connor.com"),
		Document().URL("https://example.com/document/id"),
	)

	// Display the generated structure...
	t.Log(spew.Sdump(object))

	assert.Equal(t, "Accept", object["type"])

	actor := object["actor"].(Object)
	assert.Equal(t, "john@connor.com", actor["id"])
	assert.Equal(t, "John Connor", actor["name"])
}
