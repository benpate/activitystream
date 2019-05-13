package model

import "time"

// Object data structure represents either an object or a link to an object.
// This is done because so many ActivityPub elements can be either an object or a link,
// so its easiest to model a single data type that can handle EITHER style, and then
// include validation functions that can differentiate between the two styles.
type Object struct {

	// These properties are used by both OBJECTS AND LINKS
	Context   *Context          `json:"@context,omitempty"` // Identifies the context within which the object exists or an activity was performed.
	Name      string            `json:"name,omitempty"`     // A simple, human-readable, plain-text name for the object. HTML markup MUST NOT be included.
	NameMap   map[string]string `json:"nameMap,omitEmpty"`  //
	MediaType string            `json:",omitempty"`         // Identifies the MIME media type of the value of the content property. If not specified, the content property is assumed to contain text/html content.
	Preview   *Object           `json:",omitempty"`         // Identifies an entity that provides a preview of this object.

	// These properties are used by OBJECTS ONLY
	ID           string            `json:"id,omitempty"`
	Type         []string          `json:"type,omitempty"`
	Attachment   []Object          `json:",omitempty"` // Identifies a resource attached or related to an object that potentially requires special handling. The intent is to provide a model that is at least semantically similar to attachments in email.
	AttributedTo []Object          `json:",omitempty"` // Identifies one or more entities to which this object is attributed. The attributed entities might not be Actors. For instance, an object might be attributed to the completion of another activity.
	Audience     *Object           `json:",omitempty"` // Identifies one or more entities that represent the total population of entities for which the object can considered to be relevant.
	Content      string            `json:",omitempty"` // The content or textual representation of the Object encoded as a JSON string. By default, the value of content is HTML. The mediaType property can be used in the object to indicate a different content type.
	ContentMap   map[string]string `json:",omitempty"` // The content MAY be expressed using multiple language-tagged values.
	EndTime      time.Time         `json:",omitempty"` // The date and time describing the actual or expected ending time of the object. When used with an Activity object, for instance, the endTime property specifies the moment the activity concluded or is expected to conclude.
	Generator    *Object           `json:",omitempty"` // Identifies the entity (e.g. an application) that generated the object.
	Icon         *ImageOrLink      `json:",omitempty"` // Indicates an entity that describes an icon for this object. The image should have an aspect ratio of one (horizontal) to one (vertical) and should be suitable for presentation at a small size.
	Image        *ImageOrLink      `json:",omitempty"` // Indicates an entity that describes an image for this object. Unlike the icon property, there are no aspect ratio or display size limitations assumed.
	InReplyTo    *Object           `json:",omitempty"` // Indicates one or more entities for which this object is considered a response.
	Instrument   *Object           `json:",omitempty"` // Identifies one or more objects used (or to be used) in the completion of an Activity.
	Location     *Location         `json:",omitempty"` // Indicates one or more physical or logical locations associated with the object.
	Published    time.Time         `json:",omitempty"` // The date and time at which the object was published
	Replies      *Collection       `json:",omitempty"` // Identifies a Collection containing objects considered to be responses to this object.
	StartTime    time.Time         `json:",omitempty"` // The date and time describing the actual or expected starting time of the object. When used with an Activity object, for instance, the startTime property specifies the moment the activity began or is scheduled to begin.
	Summary      string            `json:",omitempty"` // A natural language summarization of the object encoded as HTML.
	SummaryMap   map[string]string `json:",omitempty"` // Multiple language tagged summaries MAY be provided.
	Tag          []Object          `json:",omitempty"` // One or more "tags" that have been associated with an objects. A tag can be any kind of Object. The key difference between attachment and tag is that the former implies association by inclusion, while the latter implies associated by reference.
	Updated      *time.Time        `json:",omitempty"` // The date and time at which the object was updated
	URL          *Link             `json:",omitempty"` // Identifies one or more links to representations of the object
	To           []Object          `json:",omitempty"` // Identifies an entity considered to be part of the public primary audience of an Object
	BTo          []Object          `json:",omitempty"` // Identifies an Object that is part of the private primary audience of this Object.
	Cc           []Object          `json:",omitempty"` // Identifies an Object that is part of the public secondary audience of this Object.
	Bcc          []Object          `json:",omitempty"` // Identifies one or more Objects that are part of the private secondary audience of this Object.
	Duration     string            `json:",omitempty"` // When the object describes a time-bound resource, such as an audio or video, a meeting, etc, the duration property indicates the object's approximate duration. The value MUST be expressed as an xsd:duration as defined by [ xmlschema11-2], section 3.3.6 (e.g. a period of 5 seconds is represented as "PT5S").

	// The following properties are used for LINKS ONLY
	Href     string `json:"href,omitempty"`     // The target resource pointed to by a Link
	Rel      string `json:"rel,omitempty"`      // A link relation associated with a Link.
	HrefLang string `json:"hreflang,omitempty"` // Hints as to the language used by the target resource.  Must be a [BCP47] Language-Tag
	Height   int    `json:"height,omitempty"`   // Hints as to the rendering height in device-independent pixels of the linked resource.
	Width    int    `json:"width,omitemtpy"`    // Hints as to the rendering width in device-independent pixels of the linked resource.
}

/* ...These functions may end up being necessary in the future...

// MarshalJSON implements a custom JSON marshaller for this object.
// Specifically, it chooses between marshalling EITHER the Object,
// or the Link, but not both.
func (o *Object) MarshalJSON() ([]byte, error) {

	if o.Object != nil {
		return json.Marshal(o.Object)
	}

	return json.Marshal(o.Link)
}

// UnmarshalJSON implements a custom JSON marshaller for this object.
// Specifically, it chooses between unmarshalling EITHER the Object,
// or the Link, but not both.
func (o *Object) UnmarshalJSON(data []byte) error {

	structure := map[string]interface{}{}

	if err := json.Unmarshal(data, &structure); err != nil {
		return derp.New("Object.UnmarshalJSON", "Error unmarshalling JSON data", err, string(data))
	}

	return nil
}
*/
