package activitystream

// Link is an indirect, qualified reference to a resource identified by a URL. The fundamental model for links is established by [ RFC5988]. Many of the properties defined by the Activity Vocabulary allow values that are either instances of Object or Link. When a Link is used, it establishes a qualified relation connecting the subject (the containing object) to the resource identified by the href. Properties of the Link are properties of the reference as opposed to properties of the resource.
type Link struct {
	Name      string            // A simple, human-readable, plain-text name for the object. HTML markup MUST NOT be included. The name MAY be expressed using multiple language-tagged values.
	NameMap   map[string]string // A simple, human-readable, plain-text name for the object. HTML markup MUST NOT be included. The name MAY be expressed using multiple language-tagged values.
	Href      string            // URI of the targeted resource
	Hreflang  string            // BCP47 Language Tag
	Rel       string            // RFC5988 Link Relation
	MediaType string            // MIME Media Type of the referenced resource
	Height    int               // On a Link, specifies a hint as to the rendering height in device-independent pixels of the linked resource.
	Width     int               // On a Link, specifies a hint as to the rendering width in device-independent pixels of the linked resource.
	Preview   *Object           // Identifies an entity that provides a preview of this object.
}
