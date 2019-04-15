package model

import "time"

type Object struct {
	ID           string            `json:"id,omitempty"`
	Type         []string          `json:"type,omitempty"`
	Attachment   []ObjectOrLink    `json:",omitempty"` // Identifies a resource attached or related to an object that potentially requires special handling. The intent is to provide a model that is at least semantically similar to attachments in email.
	AttributedTo []ObjectOrLink    `json:",omitempty"` // Identifies one or more entities to which this object is attributed. The attributed entities might not be Actors. For instance, an object might be attributed to the completion of another activity.
	Audience     *ObjectOrLink     `json:",omitempty"` // Identifies one or more entities that represent the total population of entities for which the object can considered to be relevant.
	Content      string            `json:",omitempty"` // The content or textual representation of the Object encoded as a JSON string. By default, the value of content is HTML. The mediaType property can be used in the object to indicate a different content type.
	ContentMap   map[string]string `json:",omitempty"` // The content MAY be expressed using multiple language-tagged values.
	Context      *Object           `json:",omitempty"` // Identifies the context within which the object exists or an activity was performed.
	Name         string            `json:",omitempty"` // A simple, human-readable, plain-text name for the object. HTML markup MUST NOT be included.
	NameMap      map[string]string `json:",omitempty"` // The name MAY be expressed using multiple language-tagged values.
	EndTime      time.Time         `json:",omitempty"` // The date and time describing the actual or expected ending time of the object. When used with an Activity object, for instance, the endTime property specifies the moment the activity concluded or is expected to conclude.
	Generator    *ObjectOrLink     `json:",omitempty"` // Identifies the entity (e.g. an application) that generated the object.
	Icon         *ImageOrLink      `json:",omitempty"` // Indicates an entity that describes an icon for this object. The image should have an aspect ratio of one (horizontal) to one (vertical) and should be suitable for presentation at a small size.
	Image        *ImageOrLink      `json:",omitempty"` // Indicates an entity that describes an image for this object. Unlike the icon property, there are no aspect ratio or display size limitations assumed.
	InReplyTo    *ObjectOrLink     `json:",omitempty"` // Indicates one or more entities for which this object is considered a response.
	Instrument   *ObjectOrLink     `json:",omitempty"` // Identifies one or more objects used (or to be used) in the completion of an Activity.
	Location     *Location         `json:",omitempty"` // Indicates one or more physical or logical locations associated with the object.
	Preview      *ObjectOrLink     `json:",omitempty"` // Identifies an entity that provides a preview of this object.
	Published    time.Time         `json:",omitempty"` // The date and time at which the object was published
	Replies      *Collection       `json:",omitempty"` // Identifies a Collection containing objects considered to be responses to this object.
	StartTime    time.Time         `json:",omitempty"` // The date and time describing the actual or expected starting time of the object. When used with an Activity object, for instance, the startTime property specifies the moment the activity began or is scheduled to begin.
	Summary      string            `json:",omitempty"` // A natural language summarization of the object encoded as HTML.
	SummaryMap   map[string]string `json:",omitempty"` // Multiple language tagged summaries MAY be provided.
	Tag          []ObjectOrLink    `json:",omitempty"` // One or more "tags" that have been associated with an objects. A tag can be any kind of Object. The key difference between attachment and tag is that the former implies association by inclusion, while the latter implies associated by reference.
	Updated      *time.Time        `json:",omitempty"` // The date and time at which the object was updated
	URL          *Link             `json:",omitempty"` // Identifies one or more links to representations of the object
	To           []ObjectOrLink    `json:",omitempty"` // Identifies an entity considered to be part of the public primary audience of an Object
	BTo          []ObjectOrLink    `json:",omitempty"` // Identifies an Object that is part of the private primary audience of this Object.
	Cc           []ObjectOrLink    `json:",omitempty"` // Identifies an Object that is part of the public secondary audience of this Object.
	Bcc          []ObjectOrLink    `json:",omitempty"` // Identifies one or more Objects that are part of the private secondary audience of this Object.
	MediaType    string            `json:",omitempty"` // When used on an Object, identifies the MIME media type of the value of the content property. If not specified, the content property is assumed to contain text/html content.
	Duration     string            `json:",omitempty"` // When the object describes a time-bound resource, such as an audio or video, a meeting, etc, the duration property indicates the object's approximate duration. The value MUST be expressed as an xsd:duration as defined by [ xmlschema11-2], section 3.3.6 (e.g. a period of 5 seconds is represented as "PT5S").
}
