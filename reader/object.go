package reader

import (
	"encoding/json"
	"time"

	"github.com/benpate/derp"
	as "github.com/whisperverse/activitystream"
	"github.com/whisperverse/json-ld/reader"
)

// Object is a data structure that makes it easy(-er) to read and understand data from an ActivityStream JSON package.
// In the process, it makes irreversible changes to the original data set, so it MUST NOT be used to generate new ActivityStream
// records.  Use ObjectWriter for that purpose.
type Object struct {
	reader.JSONLD
}

/**************************
 * Constructors
 **************************/

// New converts the provided input into a fully initialized Object reference.
func New(input interface{}) Object {
	return FromJSONLD(reader.New(input))
}

// UnmarshalJSON is a shortcut for the json.Unmarshal function.  It parses the JSON input and returns a fully populated reader.Object.
func UnmarshalJSON(input string) (Object, error) {
	var result Object

	if err := json.Unmarshal([]byte(input), &result); err != nil {
		return New(""), derp.Wrap(err, "activitystream.reader.ParseJSON", "Error parsing JSON input", input)
	}

	return result, nil
}

// FromJSONLD wraps a JSONLD object as an ActivityStream.Object
func FromJSONLD(input reader.JSONLD) Object {
	return Object{JSONLD: input}
}

// SliceFromJSONLD wraps a slice of JDON-LD recoreds as a slice of ActivityStream.Objects
func SliceFromJSONLD(input []reader.JSONLD) []Object {
	result := make([]Object, len(input))

	for index, value := range input {
		result[index] = FromJSONLD(value)
	}

	return result
}

/**************************
 * ActivityPub Properties
 **************************/

// ID provides the globally unique identifier for an Object or Link
func (object Object) ID() string {
	return object.Property(as.PropertyID).AsString(as.PropertyID)
}

// Type identifies the Object or Link type.  If multiple values are present, then only the first value is returned.
func (object Object) Type() string {
	return object.Property(as.PropertyType).AsString(as.PropertyType)
}

// TypeArray lists out all types in the configured object
func (object Object) TypeArray() []string {
	return object.Property(as.PropertyType).AsSliceOfString(as.PropertyType)
}

// Actor describes one or more entities that either performed or are expected to perform the activity. Any single activity can have multiple actors. The actor MAY be specified using an indirect Link.
func (object Object) Actor() string {
	return object.Property(as.PropertyActor).AsString(as.PropertyID)
}

// ActorObject returns a fully object for the "Actor" property
func (object Object) ActorObject() Object {
	return FromJSONLD(object.Property(as.PropertyActor).AsObject(as.PropertyID))
}

// Attachment identifies a resource attached or related to an object that potentially requires special handling. The intent is to provide a model that is at least semantically similar to attachments in email.
func (object Object) Attachment() string {
	return object.Property(as.PropertyAttachment).AsString(as.PropertyURL)
}

// AttachmentObject returns a fully populated object for the "attachment" property
func (object Object) AttachmentObject() Object {
	return FromJSONLD(object.Property(as.PropertyAttachment).AsObject(as.PropertyURL))
}

// AttributedTo identifies one or more entities to which this object is attributed. The attributed entities might not be Actors. For instance, an object might be attributed to the completion of another activity.
func (object Object) AttributedTo() []string {
	return object.Property(as.PropertyAttributedTo).AsSliceOfString(as.PropertyID)
}

// AttributedToObject returns a fully populated object for the "attributedTo" property
func (object Object) AttributedToObjects() []Object {
	return SliceFromJSONLD(object.Property(as.PropertyAttributedTo).AsSliceOfObject(as.PropertyID))
}

// Audience identifies one or more entities that represent the total population of entities for which the object can considered to be relevant.
func (object Object) Audience() string {
	return object.Property(as.PropertyAudience).AsString(as.PropertyID)
}

// AudienceObject returns a fully object for the "Audience" property
func (object Object) AudienceObject() Object {
	return FromJSONLD(object.Property(as.PropertyAudience).AsObject(as.PropertyID))
}

// Bcc identifies one or more Objects that are part of the private secondary audience of this Object.
func (object Object) Bcc() []string {
	return object.Property(as.PropertyBcc).AsSliceOfString(as.PropertyID)
}

// BccObject returns a fully object for the "Bcc" property
func (object Object) BccObjects() []Object {
	return SliceFromJSONLD(object.Property(as.PropertyBcc).AsSliceOfObject(as.PropertyID))
}

// BTo identifies an Object that is part of the private primary audience of this Object.
func (object Object) BTo() []string {
	return object.Property(as.PropertyBTo).AsSliceOfString(as.PropertyID)
}

// BToObject returns a fully object for the "BTo" property
func (object Object) BToObjects() []Object {
	return SliceFromJSONLD(object.Property(as.PropertyBTo).AsSliceOfObject(as.PropertyID))
}

// Cc identifies an Object that is part of the public secondary audience of this Object.
func (object Object) Cc() []string {
	return object.Property(as.PropertyCc).AsSliceOfString(as.PropertyID)
}

// CcObject returns a fully object for the "Cc" property
func (object Object) CcObjects() []Object {
	return SliceFromJSONLD(object.Property(as.PropertyCc).AsSliceOfObject(as.PropertyID))
}

// Closed indicates that a question has been closed, and answers are no longer accepted.
func (object Object) Closed() time.Time {
	return object.Property(as.PropertyClosed).AsTime("")
}

// Content is the content or textual representation of the Object encoded as a JSON string. By default, the value of content is HTML. The mediaType property can be used in the object to indicate a different content type.  The content MAY be expressed using multiple language-tagged values.
func (object Object) Content(language ...string) string {
	return object.PropertyMap(as.PropertyContent, language...)
}

// Context Identifies the context within which the object exists or an activity was performed.  The notion of "context" used is intentionally vague. The intended function is to serve as a means of grouping objects and activities that share a common originating context or purpose. An example could be all activities relating to a common project or event.
func (object Object) Context() Object {
	return FromJSONLD(object.Property(as.PropertyContext))
}

// Duration - when the object describes a time-bound resource, such as an audio or video, a meeting, etc, the duration property indicates the object's approximate duration. The value MUST be expressed as an xsd:duration as defined by [ xmlschema11-2], section 3.3.6 (e.g. a period of 5 seconds is represented as "PT5S").
func (object Object) Duration() string {
	return object.Property(as.PropertyDuration).AsString("")
}

// EndTime represents the date and time describing the actual or expected ending time of the object. When used with an Activity object, for instance, the endTime property specifies the moment the activity concluded or is expected to conclude.
func (object Object) EndTime() time.Time {
	return object.Property(as.PropertyEndTime).AsTime("")
}

// EndTimeObject returns a fully object for the "EndTime" property
func (object Object) EndTimeObject() Object {
	return FromJSONLD(object.Property(as.PropertyEndTime).AsObject(""))
}

// Generator identifies the entity (e.g. an application) that generated the object.
func (object Object) Generator() string {
	return object.Property(as.PropertyGenerator).AsString(as.PropertyID)
}

// GeneratorObject returns a fully object for the "Generator" property
func (object Object) GeneratorObject() Object {
	return FromJSONLD(object.Property(as.PropertyGenerator).AsObject(as.PropertyID))
}

// Icon indicates an entity that describes an icon for this object. The image should have an aspect ratio of one (horizontal) to one (vertical) and should be suitable for presentation at a small size.
func (object Object) Icon() string {
	return object.Property(as.PropertyIcon).AsString(as.PropertyID)
}

// IconObject returns a fully object for the "Icon" property
func (object Object) IconObject() Object {
	return FromJSONLD(object.Property(as.PropertyIcon).AsObject(as.PropertyID))
}

// Image indicates an entity that describes an image for this object. Unlike the icon property, there are no aspect ratio or display size limitations assumed.
func (object Object) Image() string {
	return object.Property(as.PropertyImage).AsString(as.PropertyID)
}

// ImageObject returns a fully object for the "Image" property
func (object Object) ImageObject() Object {
	return FromJSONLD(object.Property(as.PropertyImage).AsObject(as.PropertyID))
}

// InReplyTo indicates one or more entities for which this object is considered a response.
func (object Object) InReplyTo() string {
	return object.Property(as.PropertyInReplyTo).AsString(as.PropertyID)
}

// InReplyToObject returns a fully object for the "InReplyTo" property
func (object Object) InReplyToObject() Object {
	return FromJSONLD(object.Property(as.PropertyInReplyTo).AsObject(as.PropertyID))
}

// Instrument identifies one or more objects used (or to be used) in the completion of an Activity.
func (object Object) Instrument() string {
	return object.Property(as.PropertyInstrument).AsString(as.PropertyID)
}

// InstrumentObject returns a fully object for the "instrument" property
func (object Object) InstrumentObject() Object {
	return FromJSONLD(object.Property(as.PropertyInstrument).AsObject(as.PropertyID))
}

// Location indicates one or more physical or logical locations associated with the object.
func (object Object) Location() string {
	return object.Property(as.PropertyLocation).AsString(as.PropertyID)
}

// LocationObject returns a fully object for the "Location" property
func (object Object) LocationObject() Object {
	return FromJSONLD(object.Property(as.PropertyLocation).AsObject(as.PropertyID))
}

// Name is a simple, human-readable, plain-text name for the object. HTML markup MUST NOT be included. The name MAY be expressed using multiple language-tagged values.
func (object Object) Name(language ...string) string {
	return object.PropertyMap(as.PropertyName, language...)
}

// Origin describes an indirect object of the activity from which the activity is directed. The precise meaning of the origin is the object of the English preposition "from". For instance, in the activity "John moved an item to List B from List A", the origin of the activity is "List A".
func (object Object) Origin() string {
	return object.Property(as.PropertyOrigin).AsString(as.PropertyID)
}

// OriginObject returns a fully object for the "Origin" property
func (object Object) OriginObject() Object {
	return FromJSONLD(object.Property(as.PropertyOrigin).AsObject(as.PropertyID))
}

// Preview identifies an entity that provides a preview of this object.
func (object Object) Preview() string {
	return object.Property(as.PropertyPreview).AsString(as.PropertyID)
}

// PreviewObject returns a fully object for the "Preview" property
func (object Object) PreviewObject() Object {
	return FromJSONLD(object.Property(as.PropertyPreview).AsObject(as.PropertyID))
}

// Published representsthe date and time at which the object was published
func (object Object) Published() time.Time {
	return object.Property(as.PropertyPublished).AsTime(as.PropertyPublished)
}

// PublishedObject returns a fully object for the "Published" property
func (object Object) PublishedObject() time.Time {
	return object.Property(as.PropertyPublished).AsTime(as.PropertyPublished)
}

/* TODO: implement collections
// Replies identifies a Collection containing objects considered to be responses to this object.
func (object Object) Replies() Collection {
	return Collection{}
}
*/

// StartTime represents the date and time describing the actual or expected starting time of the object. When used with an Activity object, for instance, the startTime property specifies the moment the activity began or is scheduled to begin
func (object Object) StartTime() time.Time {
	return object.Property(as.PropertyStartTime).AsTime(as.PropertyStartTime)
}

// StartTimeObject returns a fully object for the "StartTime" property
func (object Object) StartTimeObject() time.Time {
	return object.Property(as.PropertyStartTime).AsTime(as.PropertyStartTime)
}

// Summary is a natural language summarization of the object encoded as HTML. Multiple language tagged summaries MAY be provided.
func (object Object) Summary(language ...string) string {
	return object.PropertyMap(as.PropertySummary, language...)
}

// Tag represents one or more "tags" that have been associated with an objects. A tag can be any kind of Object. The key difference between attachment and tag is that the former implies association by inclusion, while the latter implies associated by reference.
func (object Object) Tag() string {
	return object.Property(as.PropertyTag).AsString(as.PropertyID)
}

// TagObject returns a fully object for the "Tag" property
func (object Object) TagObject() Object {
	return FromJSONLD(object.Property(as.PropertyTag).AsObject(as.PropertyID))
}

// Updated represents the date and time at which the object was updated
func (object Object) Updated() time.Time {
	return object.Property(as.PropertyUpdated).AsTime(as.PropertyUpdated)
}

// UpdatedObject returns a fully object for the "Updated" property
func (object Object) UpdatedObject() time.Time {
	return object.Property(as.PropertyUpdated).AsTime(as.PropertyUpdated)
}

// URL identifies one or more links to representations of the object
func (object Object) URL() string {
	return object.Property(as.PropertyURL).AsString(as.PropertyHREF)
}

// URLObject returns a fully object for the "URL" property
func (object Object) URLObject() Object {
	return FromJSONLD(object.Property(as.PropertyURL).AsObject(as.PropertyHREF))
}

// To identifies an entity considered to be part of the public primary audience of an Object
func (object Object) To() []string {
	return object.Property(as.PropertyTo).AsSliceOfString(as.PropertyID)
}

// ToObject returns a fully object for the "To" property
func (object Object) ToObjects() []Object {
	return SliceFromJSONLD(object.Property(as.PropertyTo).AsSliceOfObject(as.PropertyID))
}

// MediaType identifies the MIME media type of the referenced resource.  When used on an Object, identifies the MIME media type of the value of the content property. If not specified, the content property is assumed to contain text/html content.
func (object Object) MediaType() string {
	return object.Property(as.PropertyMediaType).AsString("")
}

// MediaTypeObject returns a fully object for the "MediaType" property
func (object Object) MediaTypeObject() string {
	return object.Property(as.PropertyMediaType).AsString("")
}

// HrefLang hints as to the language used by the target resource. Value MUST be a [BCP47] Language-Tag.
func (object Object) HrefLang() string {
	return object.Property(as.PropertyHrefLang).AsString("")
}

// Rel represents a link relation associated with a Link. The value MUST conform to both the [HTML5] and [RFC5988] "link relation" definitions.
func (object Object) Rel() string {
	return object.Property(as.PropertyRel).AsString("")
}

// Height specifies a hint as to the rendering height in device-independent pixels of the linked resource.
func (object Object) Height() int {
	return object.Property(as.PropertyHeight).AsInt("")
}

// Width specifies a hint as to the rendering width in device-independent pixels of the linked resource.
func (object Object) Width() int {
	return object.Property(as.PropertyWidth).AsInt("")
}

/**************************
 * Non-Standard Accessors
 **************************/

// AllRecipients collects all recipients across (to, bto, cc, bcc) into a single slice
func (object Object) AllRecipients() []string {
	result := make([]string, 0)

	result = append(result, object.To()...)
	result = append(result, object.BTo()...)
	result = append(result, object.Cc()...)
	result = append(result, object.Bcc()...)

	return result
}

/**************************
 * Other Utilities
 **************************/

// UnmarshalJSON implements the json.Unmarshaller interface
func (object *Object) UnmarshalJSON(input []byte) error {

	const location = "activitystream.reader.UnmarshalJSON"

	// First, try unmarshalling into a map (most likely)
	{
		result := make(map[string]interface{})
		if err := json.Unmarshal(input, &result); err == nil {
			object.JSONLD = reader.New(result)
			return nil
		}
	}

	// Fallthrough, try unmarshalling into a slice (less likely)
	{
		result := make([]interface{}, 0)
		if err := json.Unmarshal(input, &result); err == nil {
			object.JSONLD = reader.New(result)
			return nil
		}
	}

	// Otherwise, fail
	return derp.NewInternalError(location, "Unrecognized format", string(input))
}
