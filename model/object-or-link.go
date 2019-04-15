package model

import (
	"encoding/json"

	"github.com/benpate/derp"
)

// ObjectOrLink represents EITHER an Object OR a Link, but not both.
type ObjectOrLink struct {
	Object *Object
	Link   *Link
}

// MarshalJSON implements a custom JSON marshaller for this object.
// Specifically, it chooses between marshalling EITHER the Object,
// or the Link, but not both.
func (o *ObjectOrLink) MarshalJSON() ([]byte, error) {

	if o.Object != nil {
		return json.Marshal(o.Object)
	}

	return json.Marshal(o.Link)
}

// UnmarshalJSON implements a custom JSON marshaller for this object.
// Specifically, it chooses between unmarshalling EITHER the Object,
// or the Link, but not both.
func (o *ObjectOrLink) UnmarshalJSON(data []byte) error {

	structure := map[string]interface{}{}

	if err := json.Unmarshal(data, &structure); err != nil {
		return derp.New("ObjectOrLink.UnmarshalJSON", "Error unmarshalling JSON data", err, string(data))
	}

	return nil
}
