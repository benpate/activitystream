package activitystream

type Image struct {

	// The following fields are required for LINK objects only
	Context   *Context          `json:"@context,omitempty"`  //
	Href      string            `json:"href,omitempty"`      // The target resource pointed to by a Link
	Rel       string            `json:"rel,omitempty"`       // A link relation associated with a Link.
	MediaType string            `json:"mediaType,omitempty"` // MIME medie type of the referenced resource
	Name      string            `json:"name,omitempty"`      // A simple, human-readable, plain-text name for the object. HTML markup MUST NOT be included.
	NameMap   map[string]string `json:"nameMap,omitEmpty"`   //
	HrefLang  string            `json:"hreflang,omitempty"`  // Hints as to the language used by the target resource.  Must be a [BCP47] Language-Tag
	Height    int               `json:"height,omitempty"`    // Hints as to the rendering height in device-independent pixels of the linked resource.
	Width     int               `json:"width,omitemtpy"`     // Hints as to the rendering width in device-independent pixels of the linked resource.
	Preview   *Object           `json:"preview,omitempty"`   // Identifies an entity that provides a preview of this object

}
