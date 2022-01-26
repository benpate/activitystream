package writer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWriter(t *testing.T) {

	object := Accept(
		Person("John Connor", "en").ID("john@connor.com"),
		Document().URL("https://example.com/document/id"),
	)

	require.Equal(t, "Accept", object["type"])

	actor := object["actor"].(Object)
	require.Equal(t, "john@connor.com", actor["id"])
	require.Equal(t, "John Connor", actor["name"])
}

func TestPropertyMap(t *testing.T) {

	// Create test object
	object := Person("John Connor", "en")
	require.Equal(t, object, object.Map("", "", "en"))

	// Add mapped values
	object = object.Map("places", "Los Angeles", "en").Map("places", "Lös Angeles", "es")
	require.Equal(t, "Los Angeles", object["places"])
	require.Equal(t, 2, len(object["placesMap"].(map[string]string)))

	// Empty string does not add a value
	object = object.Map("places", "", "fr")
	require.Equal(t, 2, len(object["placesMap"].(map[string]string)))
}
