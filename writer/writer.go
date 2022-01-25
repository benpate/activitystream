package writer

import (
	"time"

	as "github.com/whisperverse/activitystream"
)

type Object map[string]interface{}

func NewObject() Object {
	return Object{}
}

// ID provides the globally unique identifier for an Object or Link
func (object Object) ID(value string) Object {
	object[as.PropertyID] = value
	return object
}

// Type identifies the Object or Link type.  If multiple values are present, then only the first value is returned.
func (object Object) Type(value interface{}) Object {
	return object.Property(as.PropertyType, value)
}

// Actor describes one or more entities that either performed or are expected to perform the activity. Any single activity can have multiple actors. The actor MAY be specified using an indirect Link.
func (object Object) Actor(value interface{}) Object {
	return object.Property(as.PropertyActor, value)
}

// Attachment identifies a resource attached or related to an object that potentially requires special handling. The intent is to provide a model that is at least semantically similar to attachments in email.
func (object Object) Attachment(value interface{}) Object {
	return object.Property(as.PropertyAttachment, value)
}

// AttributedTo identifies one or more entities to which this object is attributed. The attributed entities might not be Actors. For instance, an object might be attributed to the completion of another activity.
func (object Object) AttributedTo(value interface{}) Object {
	return object.Property(as.PropertyAttributedTo, value)
}

// Audience identifies one or more entities that represent the total population of entities for which the object can considered to be relevant.
func (object Object) Audience(value interface{}) Object {
	return object.Property(as.PropertyAudience, value)
}

// Bcc identifies one or more Objects that are part of the private secondary audience of this Object.
func (object Object) Bcc(value interface{}) Object {
	return object.Property(as.PropertyBcc, value)
}

// BTo identifies an Object that is part of the private primary audience of this Object.
func (object Object) BTo(value interface{}) Object {
	return object.Property(as.PropertyBTo, value)
}

// Cc identifies an Object that is part of the aslic secondary audience of this Object.
func (object Object) Cc(value interface{}) Object {
	return object.Property(as.PropertyCc, value)
}

// Closed indicates that a question has been closed, and answers are no longer accepted.
func (object Object) Closed(value time.Time) Object {
	return object.Property(as.PropertyClosed, value)
}

// Content is the content or textual representation of the Object encoded as a JSON string. By default, the value of content is HTML. The mediaType property can be used in the object to indicate a different content type.  The content MAY be expressed using multiple language-tagged values.
func (object Object) Content(value string, language string) Object {
	return object.Map(as.PropertyContent, value, language)
}

// Context Identifies the context within which the object exists or an activity was performed.  The notion of "context" used is intentionally vague. The intended function is to serve as a means of grouping objects and activities that share a common originating context or purpose. An example could be all activities relating to a common project or event.
func (object Object) Context(value interface{}) Object {
	// TODO: incomplete
	return object
}

// Duration - when the object describes a time-bound resource, such as an audio or video, a meeting, etc, the duration property indicates the object's approximate duration. The value MUST be expressed as an xsd:duration as defined by [ xmlschema11-2], section 3.3.6 (e.g. a period of 5 seconds is represented as "PT5S").
func (object Object) Duration(value time.Duration) Object {
	return object.SimpleValue(as.PropertyDuration, value)
}

// EndTime represents the date and time describing the actual or expected ending time of the object. When used with an Activity object, for instance, the endTime property specifies the moment the activity concluded or is expected to conclude.
func (object Object) EndTime(value time.Time) Object {
	return object.Property(as.PropertyEndTime, value)
}

// Generator identifies the entity (e.g. an application) that generated the object.
func (object Object) Generator(value interface{}) Object {
	return object.Property(as.PropertyGenerator, value)
}

// Icon indicates an entity that describes an icon for this object. The image should have an aspect ratio of one (horizontal) to one (vertical) and should be suitable for presentation at a small size.
func (object Object) Icon(value interface{}) Object {
	return object.Property(as.PropertyIcon, value)
}

// Image indicates an entity that describes an image for this object. Unlike the icon property, there are no aspect ratio or display size limitations assumed.
func (object Object) Image(value interface{}) Object {
	return object.Property(as.PropertyImage, value)
}

// InReplyTo indicates one or more entities for which this object is considered a response.
func (object Object) InReplyTo(value interface{}) Object {
	return object.Property(as.PropertyInReplyTo, value)
}

// Instrument identifies one or more objects used (or to be used) in the completion of an Activity.
func (object Object) Instrument(value interface{}) Object {
	return object.Property(as.PropertyInstrument, value)
}

// Location indicates one or more physical or logical locations associated with the object.
func (object Object) Location(value interface{}) Object {
	return object.Property(as.PropertyLocation, value)
}

// Name is a simple, human-readable, plain-text name for the object. HTML markup MUST NOT be included. The name MAY be expressed using multiple language-tagged values.
func (object Object) Name(value string, language string) Object {
	return object.Map(as.PropertyName, value, language)
}

// Object
func (object Object) Object(value interface{}) Object {
	return object.Property(as.PropertyObject, value)
}

// Origin describes an indirect object of the activity from which the activity is directed. The precise meaning of the origin is the object of the English preposition "from". For instance, in the activity "John moved an item to List B from List A", the origin of the activity is "List A".
func (object Object) Origin(value interface{}) Object {
	return object.Property(as.PropertyOrigin, value)
}

// Preview identifies an entity that provides a preview of this object.
func (object Object) Preview(value interface{}) Object {
	return object.Property(as.PropertyPreview, value)
}

// Published representsthe date and time at which the object was aslished
func (object Object) Published(value time.Time) Object {
	return object.Property(as.PropertyPublished, value)
}

// Replies identifies a Collection containing objects considered to be responses to this object.
/*func (object Object) Replies(value interface{}) Collection {
	return Collection{}
}*/

// StartTime represents the date and time describing the actual or expected starting time of the object. When used with an Activity object, for instance, the startTime property specifies the moment the activity began or is scheduled to begin
func (object Object) StartTime(value time.Time) Object {
	return object.Property(as.PropertyStartTime, value)
}

// Summary is a natural language summarization of the object encoded as HTML. Multiple language tagged summaries MAY be provided.
func (object Object) Summary(value string, language string) Object {
	return object.Map(as.PropertySummary, value, language)
}

// Tag represents one or more "tags" that have been associated with an objects. A tag can be any kind of Object. The key difference between attachment and tag is that the former implies association by inclusion, while the latter implies associated by reference.
func (object Object) Tag(value interface{}) Object {
	return object.Property(as.PropertyTag, value)
}

// Target
func (object Object) Target(value interface{}) Object {
	return object.Property(as.PropertyTarget, value)
}

// To identifies an entity considered to be part of the aslic primary audience of an Object
func (object Object) To(value interface{}) Object {
	return object.Property(as.PropertyTo, value)
}

// Updated represents the date and time at which the object was updated
func (object Object) Updated(value time.Time) Object {
	return object.Property(as.PropertyUpdated, value)
}

// URL identifies one or more links to representations of the object
func (object Object) URL(value interface{}) Object {
	return object.Property(as.PropertyURL, value)
}

// MediaType identifies the MIME media type of the referenced resource.  When used on an Object, identifies the MIME media type of the value of the content property. If not specified, the content property is assumed to contain text/html content.
func (object Object) MediaType(value interface{}) Object {
	return object.Property(as.PropertyMediaType, value)
}

// HrefLang hints as to the language used by the target resource. Value MUST be a [BCP47] Language-Tag.
func (object Object) HrefLang(value interface{}) Object {
	return object.Property(as.PropertyHrefLang, value)
}

// Rel represents a link relation associated with a Link. The value MUST conform to both the [HTML5] and [RFC5988] "link relation" definitions.
func (object Object) Rel(value interface{}) Object {
	return object.Property(as.PropertyRel, value)
}

// Height specifies a hint as to the rendering height in device-independent pixels of the linked resource.
func (object Object) Height(value int64) Object {
	return object.SimpleValue(as.PropertyHeight, value)
}

// Width specifies a hint as to the rendering width in device-independent pixels of the linked resource.
func (object Object) Width(value int64) Object {
	return object.SimpleValue(as.PropertyWidth, value)
}

// Property sets the value of a property to whatever value is provided
func (object Object) Property(property string, value interface{}) Object {

	// If the map does not have this key in it, then insert.
	if _, ok := object[property]; !ok {
		object[property] = value
		return object
	}

	switch current := object[property].(type) {

	case []interface{}:
		// If we already have an array, then append to it.
		object[property] = append(current, value)

	default:
		// Otherwise, make a new array, and append to it.
		// object[key] = []interface{current, value}
	}

	return object
}

// Map updates values that may/may-not be multi-language maps.
func (object Object) Map(property string, value string, language string) Object {

	// If we don't already have a default property, then add it now.
	if _, ok := object[property]; !ok {
		object[property] = value
	}

	// If we don't already have a propertyMap, then add it now
	propertyMap := property + "Map"
	if _, ok := object[propertyMap]; !ok {
		object[propertyMap] = map[string]string{}
	}

	// Safely set the value of the propertyMap
	if propertyMap, ok := object[propertyMap].(map[string]string); ok {
		propertyMap[language] = value
	}

	// Success!
	return object
}

// SimpleValue assigns a value to a property with no other shenanigans.
func (object Object) SimpleValue(property string, value interface{}) Object {
	object[property] = value
	return object
}
