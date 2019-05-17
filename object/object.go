package object

import "time"

type Object struct {
	ID           string            `json:"id,omitempty"`           // Provides the globally unique identifier for an Object or Link.
	Type         []string          `json:"type,omitempty"`         //  	Identifies the Object or Link type. Multiple values may be specified.
	AttributedTo *Object           `json:"attributedTo,omitempty"` // Identifies one or more entities to which this object is attributed. The attributed entities might not be Actors. For instance, an object might be attributed to the completion of another activity.
	Name         string            `json:"name,omitempty"`         // A simple, human-readable, plain-text name for the object. HTML markup MUST NOT be included.
	NameMap      map[string]string `json:"nameMap,omitEmpty"`      // Multiple language tagged names MAY be provided.

	Context    *Context          `json:"@context,omitempty"`   // Identifies the context within which the object exists or an activity was performed.
	MediaType  string            `json:"mediaType,omitempty"`  // Identifies the MIME media type of the value of the content property. If not specified, the content property is assumed to contain text/html content.
	Preview    *Object           `json:"preview,omitempty"`    // Identifies an entity that provides a preview of this object.
	Attachment *Object           `json:"attachment,omitempty"` // Identifies a resource attached or related to an object that potentially requires special handling. The intent is to provide a model that is at least semantically similar to attachments in email.
	Audience   *Object           `json:"audience,omitempty"`   // Identifies one or more entities that represent the total population of entities for which the object can considered to be relevant.
	Content    string            `json:"content,omitempty"`    // The content or textual representation of the Object encoded as a JSON string. By default, the value of content is HTML. The mediaType property can be used in the object to indicate a different content type.
	ContentMap map[string]string `json:"contentMap,omitempty"` // The content MAY be expressed using multiple language-tagged values.
	EndTime    time.Time         `json:"endTime,omitempty"`    // The date and time describing the actual or expected ending time of the object. When used with an Activity object, for instance, the endTime property specifies the moment the activity concluded or is expected to conclude.
	Generator  *Object           `json:"generator,omitempty"`  // Identifies the entity (e.g. an application) that generated the object.
	Image      *Object           `json:"image,omitempty"`      // Indicates an entity that describes an image for this object. Unlike the icon property, there are no aspect ratio or display size limitations assumed.
	InReplyTo  *Object           `json:"inReplyTo,omitempty"`  // Indicates one or more entities for which this object is considered a response.
	Location   *Object           `json:"location,omitempty"`   // Indicates one or more physical or logical locations associated with the object.
	Published  time.Time         `json:"published,omitempty"`  // The date and time at which the object was published
	Replies    *Collection       `json:"repllies,omitempty"`   // Identifies a Collection containing objects considered to be responses to this object.
	StartTime  time.Time         `json:"startTime,omitempty"`  // The date and time describing the actual or expected starting time of the object. When used with an Activity object, for instance, the startTime property specifies the moment the activity began or is scheduled to begin.
	Summary    string            `json:"summary,omitempty"`    // A natural language summarization of the object encoded as HTML.
	SummaryMap map[string]string `json:"summaryMap,omitempty"` // Multiple language tagged summaries MAY be provided.
	Tag        []Object          `json:"tag,omitempty"`        // One or more "tags" that have been associated with an objects. A tag can be any kind of Object. The key difference between attachment and tag is that the former implies association by inclusion, while the latter implies associated by reference.
	Updated    *time.Time        `json:"updated,omitempty"`    // The date and time at which the object was updated
	URL        *Link             `json:"url,omitempty"`        // Identifies one or more links to representations of the object
	To         *Object           `json:"to,omitempty"`         // Identifies an entity considered to be part of the public primary audience of an Object
	BTo        *Object           `json:"bto,omitempty"`        // Identifies an Object that is part of the private primary audience of this Object.
	Cc         *Object           `json:"cc,omitempty"`         // Identifies an Object that is part of the public secondary audience of this Object.
	Bcc        *Object           `json:"bcc,omitempty"`        // Identifies one or more Objects that are part of the private secondary audience of this Object.
	Duration   string            `json:"duration,omitempty"`   // When the object describes a time-bound resource, such as an audio or video, a meeting, etc, the duration property indicates the object's approximate duration. The value MUST be expressed as an xsd:duration as defined by [ xmlschema11-2], section 3.3.6 (e.g. a period of 5 seconds is represented as "PT5S").

	Icon *Object `json:"icon,omitempty"` // Indicates an entity that describes an icon for this object. The image should have an aspect ratio of one (horizontal) to one (vertical) and should be suitable for presentation at a small size.

}
