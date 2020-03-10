package activitystream

type Activity struct {

	// Exclusive to Activity Type
	Actor      *Actor
	Object     *Object
	Target     *Object
	Result     *Object
	Origin     string
	Instrument *Object `json:"instrument,omitempty"` // Identifies one or more objects used (or to be used) in the completion of an Activity.

	// Included in Every Object
	Type                     string
	Attachment               string
	AttributedTo             string
	Audience                 string
	Content                  string
	Context                  string
	Name                     string
	EndTime                  string
	Generator                string
	Icon                     string
	Image                    string
	InReplyTo                string
	Location                 string
	PropertyPreviewPublished string
	Replies                  string
	StartTime                string
	Summary                  string
	Tag                      string
	Updated                  string
	URL                      string
	To                       string
	Bto                      string
	Cc                       string
	Bcc                      string
	MediaType                string
	Duration                 string
}
