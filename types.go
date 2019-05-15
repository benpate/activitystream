package activitystream

// These types are defined by the W3C on https://www.w3.org/TR/activitystreams-vocabulary/#object-types

/**** Object Types ****/

const (

	//ObjectTypeArticle represents any kind of multi-paragraph written work.
	ObjectTypeArticle = "Article"

	//ObjectTypeAudio represents an audio document of any kind.
	ObjectTypeAudio = "Audio"

	//ObjectTypeDocument represents a document of any kind.
	ObjectTypeDocument = "Document"

	//ObjectTypeEvent represents any kind of event.
	ObjectTypeEvent = "Event"

	//ObjectTypeImage represents an image document of any kind
	ObjectTypeImage = "Image"

	//ObjectTypeNote represents a short written work typically less than a single paragraph in length.
	ObjectTypeNote = "Note"

	//ObjectTypePage represents a Web Page.
	ObjectTypePage = "Page"

	//ObjectTypePlace represents a logical or physical location. See 5.3 Representing Places for additional information.
	ObjectTypePlace = "Place"

	//ObjectTypeProfile is a content object that describes another Object, typically used to describe Actor Type objects. The describes property is used to reference the object being described by the profile.
	ObjectTypeProfile = "Profile"

	//ObjectTypeRelationship describes a relationship between two individuals. The subject and object properties are used to identify the connected individuals.
	ObjectTypeRelationship = "Relationship"

	//ObjectTypeTombstone represents a content object that has been deleted. It can be used in Collections to signify that there used to be an object at this position, but it has been deleted.
	ObjectTypeTombstone = "Tombstone"

	//ObjectTypeVideo represents a video document of any kind.
	ObjectTypeVideo = "Video"
)

/**** Link Types ****/

const (
	// LinkTypeMention is a specialized Link that represents an @mention.
	LinkTypeMention = "Mention"
)
