package reader

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/benpate/activitystream/vocabulary"
	"github.com/benpate/derp"
)

// Object is a data structure that makes it easy(-er) to read and understand data from an ActivityStream JSON package.
// In the process, it makes irreversible changes to the original data set, so it MUST NOT be used to generate new ActivityStream
// records.  Use ObjectWriter for that purpose.
type Object struct {
	currentRecord map[string]interface{}
	nextRecords   []map[string]interface{}
}

// NewObject converts the provided input into a fully initialized Object reference, or returns an error trying.
func NewObject(input interface{}) (*Object, error) {

	object := Object{
		currentRecord: map[string]interface{}{},
		nextRecords:   []map[string]interface{}{},
	}

	switch input := input.(type) {
	case string:
		if err := json.Unmarshal([]byte(input), &object.currentRecord); err != nil {
			return nil, derp.New(http.StatusBadRequest, "activitystream.NewObject", "Unable to unmarshal JSON", err, input)
		}

	default:
		return nil, derp.New(http.StatusBadRequest, "activitystream.NewObject", "Unrecognized input type", input)
	}

	return &object, nil
}

// ID provides the globally unique identifier for an Object or Link
func (object Object) ID() string {
	result, _ := object.ParseString(vocabulary.PropertyID)
	return result
}

// Type identifies the Object or Link type.  Multiple values may be supported
func (object Object) Type() []string {
	result, _ := object.ParseArray(vocabulary.PropertyType)
	return result
}

// Actor describes one or more entities that either performed or are expected to perform the activity. Any single activity can have multiple actors. The actor MAY be specified using an indirect Link.
func (object Object) Actor() *Object {
	actor, _ := object.ParseObject(vocabulary.PropertyActor)
	return actor
}

// ActorOk returns the "actor" property of the current object, along with a boolean that identifies if the property is present or not
func (object Object) ActorOk() (*Object, bool) {
	return object.ParseObject(vocabulary.PropertyActor)
}

// Attachment identifies a resource attached or related to an object that potentially requires special handling. The intent is to provide a model that is at least semantically similar to attachments in email.
func (object Object) Attachment() *Object {
	result, _ := object.ParseObject(vocabulary.PropertyAttachment)
	return result
}

// AttachmentOk returns the "attachment" property of the current object, along with a boolean that identifies if the property is present or not
func (object Object) AttachmentOk() (*Object, bool) {
	return object.ParseObject(vocabulary.PropertyAttachment)
}

// AttributedTo identifies one or more entities to which this object is attributed. The attributed entities might not be Actors. For instance, an object might be attributed to the completion of another activity.
func (object Object) AttributedTo() *Object {
	result, _ := object.ParseObject(vocabulary.PropertyAttributedTo)
	return result
}

// AttributedToOk returns the "attributedTo" property of the current object, along with a boolean that identifies if the property is present or not
func (object Object) AttributedToOk() (*Object, bool) {
	return object.ParseObject(vocabulary.PropertyAttributedTo)
}

// Audience identifies one or more entities that represent the total population of entities for which the object can considered to be relevant.
func (object Object) Audience() *Object {
	result, _ := object.ParseObject(vocabulary.PropertyAudience)
	return result
}

// AudienceOk returns the "audience" property of the current object, along with a boolean that identifies if the property is present or not
func (object Object) AudienceOk() (*Object, bool) {
	return object.ParseObject(vocabulary.PropertyAudience)
}

// Bcc identifies one or more Objects that are part of the private secondary audience of this Object.
func (object Object) Bcc() *Object {
	result, _ := object.ParseObject(vocabulary.PropertyBcc)
	return result
}

// BccOk returns the "bcc" property of the current object, along with a boolean that identifies if the property is present or not
func (object Object) BccOk() (*Object, bool) {
	return object.ParseObject(vocabulary.PropertyBcc)
}

// BTo identifies an Object that is part of the private primary audience of this Object.
func (object Object) BTo() *Object {
	result, _ := object.ParseObject(vocabulary.PropertyBTo)
	return result
}

// BToOk returns the "bto" property of the current object, along with a boolean that identifies if the property is present or not
func (object Object) BToOk() (*Object, bool) {
	return object.ParseObject(vocabulary.PropertyBTo)
}

// Cc identifies an Object that is part of the public secondary audience of this Object.
func (object Object) Cc() *Object {
	result, _ := object.ParseObject(vocabulary.PropertyCc)
	return result
}

// CcOk returns the "cc" property of the current object, along with a boolean that identifies if the property is present or not
func (object Object) CcOk() (*Object, bool) {
	return object.ParseObject(vocabulary.PropertyCc)
}

// Closed indicates that a question has been closed, and answers are no longer accepted.
func (object Object) Closed() time.Time {
	result, _ := object.ParseTime(vocabulary.PropertyClosed)
	return result
}

// ClosedOk returns the "closed" property of the current object, along with a boolean that identifies if the property is present or not
func (object Object) ClosedOk() (time.Time, bool) {
	return object.ParseTime(vocabulary.PropertyClosed)
}

//  Content is the content or textual representation of the Object encoded as a JSON string. By default, the value of content is HTML. The mediaType property can be used in the object to indicate a different content type.  The content MAY be expressed using multiple language-tagged values.
func (object Object) Content(language string) string {
	result, _ := object.ParseString(vocabulary.PropertyContent)
	return result
}

// ContentOk returns the "content" property of the current object, along with a boolean that identifies if the property is present or not
func (object Object) ContentOk(language string) (string, bool) {
	return object.ParseString(vocabulary.PropertyContent)
}

// Context Identifies the context within which the object exists or an activity was performed.  The notion of "context" used is intentionally vague. The intended function is to serve as a means of grouping objects and activities that share a common originating context or purpose. An example could be all activities relating to a common project or event.
func (object Object) Context() Context {
	result, _ := object.ParseContext()
	return result
}

// ContextOk returns the "context" property of the current object, along with a boolean that identifies if the property is present or not
func (object Object) ContextOk() (Context, bool) {
	return object.ParseContext()
}

//  Duration: When the object describes a time-bound resource, such as an audio or video, a meeting, etc, the duration property indicates the object's approximate duration. The value MUST be expressed as an xsd:duration as defined by [ xmlschema11-2], section 3.3.6 (e.g. a period of 5 seconds is represented as "PT5S").
func (object Object) Duration() time.Duration {
	result, _ := object.ParseDuration(vocabulary.PropertyDuration)
	return result
}

// DurationOk returns the "duration" property of the current object, along with a boolean that identifies if the property is present or not
func (object Object) DurationOk() (time.Duration, bool) {
	return object.ParseDuration(vocabulary.PropertyDuration)
}

// EndTime represents the date and time describing the actual or expected ending time of the object. When used with an Activity object, for instance, the endTime property specifies the moment the activity concluded or is expected to conclude.
func (object Object) EndTime() time.Time {
	result, _ := object.ParseTime(vocabulary.PropertyEndTime)
	return result
}

// EndTimeOk returns the "endTime" property of the current object, along with a boolean that identifies if the property is present or not
func (object Object) EndTimeOk() (time.Time, bool) {
	return object.ParseTime(vocabulary.PropertyEndTime)
}

// Generator identifies the entity (e.g. an application) that generated the object.
func (object Object) Generator() *Object {
	result, _ := object.ParseObject(vocabulary.PropertyGenerator)
	return result
}

func (object Object) GeneratorOk() (*Object, bool) {
	return object.ParseObject(vocabulary.PropertyGenerator)
}

// Icon indicates an entity that describes an icon for this object. The image should have an aspect ratio of one (horizontal) to one (vertical) and should be suitable for presentation at a small size.
func (object Object) Icon() *Object {
	result, _ := object.ParseObject(vocabulary.PropertyIcon)
	return result
}

func (object Object) IconOk() (*Object, bool) {
	return object.ParseObject(vocabulary.PropertyIcon)
}

// Image indicates an entity that describes an image for this object. Unlike the icon property, there are no aspect ratio or display size limitations assumed.
func (object Object) Image() *Object {
	result, _ := object.ParseObject(vocabulary.PropertyImage)
	return result
}

func (object Object) ImageOk() (*Object, bool) {
	return object.ParseObject(vocabulary.PropertyImage)
}

// InReplyTo indicates one or more entities for which this object is considered a response.
func (object Object) InReplyTo() *Object {
	result, _ := object.ParseObject(vocabulary.PropertyInReplyTo)
	return result
}

func (object Object) InReplyToOk() (*Object, bool) {
	return object.ParseObject(vocabulary.PropertyInReplyTo)
}

// Instrument identifies one or more objects used (or to be used) in the completion of an Activity.
func (object Object) Instrument() *Object {
	result, _ := object.ParseObject(vocabulary.PropertyInstrument)
	return result
}

func (object Object) InstrumentOk() (*Object, bool) {
	return object.ParseObject(vocabulary.PropertyInstrument)
}

// Indicates one or more physical or logical locations associated with the object.
func (object Object) Location() *Object {
	result, _ := object.ParseObject(vocabulary.PropertyLocation)
	return result
}

func (object Object) LocationOk() (*Object, bool) {
	return object.ParseObject(vocabulary.PropertyLocation)
}

// Name is a simple, human-readable, plain-text name for the object. HTML markup MUST NOT be included. The name MAY be expressed using multiple language-tagged values.
func (object Object) Name(language string) string {
	result, _ := object.ParseMap(vocabulary.PropertyName, language)
	return result
}

// NameOk returns the "name" property of the current object, along with a boolean that identifies if the property is present or not
func (object Object) NameOk(language string) (string, bool) {
	return object.ParseMap(vocabulary.PropertyName, language)
}

// Origin describes an indirect object of the activity from which the activity is directed. The precise meaning of the origin is the object of the English preposition "from". For instance, in the activity "John moved an item to List B from List A", the origin of the activity is "List A".
func (object Object) Origin() *Object {
	result, _ := object.ParseObject(vocabulary.PropertyOrigin)
	return result
}

func (object Object) OriginOk() (*Object, bool) {
	return object.ParseObject(vocabulary.PropertyOrigin)
}

// Preview identifies an entity that provides a preview of this object.
func (object Object) Preview() *Object {
	result, _ := object.ParseObject(vocabulary.PropertyPreview)
	return result
}

func (object Object) PreviewOk() (*Object, bool) {
	return object.ParseObject(vocabulary.PropertyPreview)
}

// Published representsthe date and time at which the object was published
func (object Object) Published() time.Time {
	result, _ := object.ParseTime(vocabulary.PropertyPublished)
	return result
}

func (object Object) PublishedOk() (time.Time, bool) {
	return object.ParseTime(vocabulary.PropertyPublished)
}

// Replies identifies a Collection containing objects considered to be responses to this object.
func (object Object) Replies() Collection {
	return Collection{}
}

// StartTime represents the date and time describing the actual or expected starting time of the object. When used with an Activity object, for instance, the startTime property specifies the moment the activity began or is scheduled to begin
func (object Object) StartTime() time.Time {
	result, _ := object.ParseTime(vocabulary.PropertyStartTime)
	return result
}

func (object Object) StartTimeOk() (time.Time, bool) {
	return object.ParseTime(vocabulary.PropertyStartTime)
}

// Summary is a natural language summarization of the object encoded as HTML. Multiple language tagged summaries MAY be provided.
func (object Object) Summary(language string) string {
	result, _ := object.ParseMap(vocabulary.PropertySummary, language)
	return result
}

func (object Object) SummaryOk(language string) (string, bool) {
	return object.ParseMap(vocabulary.PropertySummary, language)
}

// Tag represents one or more "tags" that have been associated with an objects. A tag can be any kind of Object. The key difference between attachment and tag is that the former implies association by inclusion, while the latter implies associated by reference.
func (object Object) Tag() *Object {
	result, _ := object.ParseObject(vocabulary.PropertyTag)
	return result
}

func (object Object) TagOk() (*Object, bool) {
	return object.ParseObject(vocabulary.PropertyTag)
}

// Updated represents the date and time at which the object was updated
func (object Object) Updated() time.Time {
	result, _ := object.ParseTime(vocabulary.PropertyUpdated)
	return result
}

func (object Object) UpdatedOk() (time.Time, bool) {
	return object.ParseTime(vocabulary.PropertyUpdated)
}

// URL identifies one or more links to representations of the object
func (object Object) URL() *Object {
	result, _ := object.ParseObject(vocabulary.PropertyURL)
	return result
}

func (object Object) URLOk() (*Object, bool) {
	return object.ParseObject(vocabulary.PropertyURL)
}

// To identifies an entity considered to be part of the public primary audience of an Object
func (object Object) To() *Object {
	result, _ := object.ParseObject(vocabulary.PropertyTo)
	return result
}

func (object Object) ToOk() (*Object, bool) {
	return object.ParseObject(vocabulary.PropertyTo)
}

//  When used on a Link, identifies the MIME media type of the referenced resource.  When used on an Object, identifies the MIME media type of the value of the content property. If not specified, the content property is assumed to contain text/html content.
func (object Object) MediaType() string {
	result, _ := object.ParseString(vocabulary.PropertyMediaType)
	return result
}

func (object Object) MediaTypeOk() (string, bool) {
	return object.ParseString(vocabulary.PropertyMediaType)
}

// HrefLang hints as to the language used by the target resource. Value MUST be a [BCP47] Language-Tag.
func (object Object) HrefLang() string {
	result, _ := object.ParseString(vocabulary.PropertyHrefLang)
	return result
}

func (object Object) HrefLangOk() (string, bool) {
	return object.ParseString(vocabulary.PropertyHrefLang)
}

// Rel represents a link relation associated with a Link. The value MUST conform to both the [HTML5] and [RFC5988] "link relation" definitions.
func (object Object) Rel() string {
	result, _ := object.ParseString(vocabulary.PropertyRel)
	return result
}

func (object Object) RelOk() (string, bool) {
	return object.ParseString(vocabulary.PropertyRel)
}

// Height specifies a hint as to the rendering height in device-independent pixels of the linked resource.
func (object Object) Height() int {
	result, _ := object.ParseInt(vocabulary.PropertyHeight)
	return result
}

func (object Object) HeightOk() (int, bool) {
	return object.ParseInt(vocabulary.PropertyHeight)
}

// Width specifies a hint as to the rendering width in device-independent pixels of the linked resource.
func (object Object) Width() int {
	result, _ := object.ParseInt(vocabulary.PropertyWidth)
	return result
}

func (object Object) WidthOk() (int, bool) {
	return object.ParseInt(vocabulary.PropertyWidth)
}

func (object Object) ParseContext() (Context, bool) {
	return Context{}, false
}

func (object Object) ParseObject(property string) (*Object, bool) {
	return nil, false
}

func (object Object) ParseString(property string) (string, bool) {
	return "", false
}

func (object Object) ParseArray(property string) ([]string, bool) {
	return []string{}, false
}

func (object Object) ParseInt(property string) (int, bool) {
	return 0, false
}

func (object Object) ParseMap(property string, language string) (string, bool) {
	return "", false
}

func (object Object) ParseTime(property string) (time.Time, bool) {
	return time.Time{}, false
}

func (object Object) ParseDuration(property string) (time.Duration, bool) {
	return time.Duration(0), false
}

func (object *Object) UnmarshallJSON(input []byte) error {

	var rawValue interface{}

	if err := json.Unmarshal(input, rawValue); err != nil {
		return derp.New(http.StatusBadRequest, "activitystream.object", "Unable to unmarshal JSON", err, input)
	}

	return nil
}
