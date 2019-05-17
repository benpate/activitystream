package object

type Link struct {
	ID           string            `json:"id,omitempty"`           // Provides the globally unique identifier for an Object or Link.
	Type         []string          `json:"type,omitempty"`         //  	Identifies the Object or Link type. Multiple values may be specified.
	Name         string            `json:"name,omitempty"`         // A simple, human-readable, plain-text name for the object. HTML markup MUST NOT be included.
	NameMap      map[string]string `json:"nameMap,omitEmpty"`      // Multiple language tagged names MAY be provided.
	AttributedTo *Object           `json:"attributedTo,omitempty"` // Identifies one or more entities to which this object is attributed. The attributed entities might not be Actors. For instance, an object might be attributed to the completion of another activity.
	Preview      *Object           `json:"preview,omitempt"`       // Identifies an entity that provides a preview of this object.
}
