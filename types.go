package model

// These types are defined by the W3C on https://www.w3.org/TR/activitystreams-vocabulary/#object-types

/**** Activity Types ****/

const (
	// ActivityTypeAccept indicates that the actor accepts the object. The target property can be used in certain circumstances to indicate the context into which the object has been accepted.
	ActivityTypeAccept = "Accept"

	// ActivityTypeAdd indicates that the actor has added the object to the target. If the target property is not explicitly specified, the target would need to be determined implicitly by context. The origin can be used to identify the context from which the object originated.
	ActivityTypeAdd = "Add"

	// ActivityTypeAnnounce indicates that the actor is calling the target's attention the object.
	ActivityTypeAnnounce = "Announce"

	// ActivityTypeArrive is an IntransitiveActivity that indicates that the actor has arrived at the location. The origin can be used to identify the context from which the actor originated. The target typically has no defined meaning.
	ActivityTypeArrive = "Arrive"

	// ActivityTypeBlock indicates that the actor is blocking the object. Blocking is a stronger form of Ignore. The typical use is to support social systems that allow one user to block activities or content of other users. The target and origin typically have no defined meaning.
	ActivityTypeBlock = "Block"

	// ActivityTypeCreate indicates that the actor has created the object.
	ActivityTypeCreate = "Create"

	// ActivityTypeDelete indicates that the actor has deleted the object. If specified, the origin indicates the context from which the object was deleted.
	ActivityTypeDelete = "Delete"

	// ActivityTypeDislike indicates that the actor dislikes the object.
	ActivityTypeDislike = "Dislike"

	// ActivityTypeFlag indicates that the actor is "flagging" the object. Flagging is defined in the sense common to many social platforms as reporting content as being inappropriate for any number of reasons.
	ActivityTypeFlag = "Flag"

	// ActivityTypeFollow indicates that the actor is "following" the object. Following is defined in the sense typically used within Social systems in which the actor is interested in any activity performed by or on the object. The target and origin typically have no defined meaning.
	ActivityTypeFollow = "Follow"

	// ActivityTypeIgnore indicates that the actor is ignoring the object. The target and origin typically have no defined meaning.
	ActivityTypeIgnore = "Ignore"

	// ActivityTypeInvite is a specialization of Offer in which the actor is extending an invitation for the object to the target.
	ActivityTypeInvite = "Invite"

	// ActivityTypeJoin indicates that the actor has joined the object. The target and origin typically have no defined meaning.
	ActivityTypeJoin = "Join"

	// ActivityTypeLeave indicates that the actor has left the object. The target and origin typically have no meaning.
	ActivityTypeLeave = "Leave"

	// ActivityTypeLike indicates that the actor likes, recommends or endorses the object. The target and origin typically have no defined meaning.
	ActivityTypeLike = "Like"

	// ActivityTypeListen indicates that the actor has listened to the object.
	ActivityTypeListen = "Listen"

	// ActivityTypeMove indicates that the actor has moved object from origin to target. If the origin or target are not specified, either can be determined by context.
	ActivityTypeMove = "Move"

	// ActivityTypeOffer indicates that the actor is offering the object. If specified, the target indicates the entity to which the object is being offered.
	ActivityTypeOffer = "Offer"

	// ActivityTypeQuestion represents a question being asked. Question objects are an extension of IntransitiveActivity. That is, the Question object is an Activity, but the direct object is the question itself and therefore it would not contain an object property.
	ActivityTypeQuestion = "Question"

	// ActivityTypeReject indicates that the actor is rejecting the object. The target and origin typically have no defined meaning.
	ActivityTypeReject = "Reject"

	// ActivityTypeRead indicates that the actor has read the object.
	ActivityTypeRead = "Read"

	// ActivityTypeRemove indicates that the actor is removing the object. If specified, the origin indicates the context from which the object is being removed.
	ActivityTypeRemove = "Remove"

	// ActivityTypeTentativeReject is a specialization of Reject in which the rejection is considered tentative.
	ActivityTypeTentativeReject = "TentativeReject"

	// ActivityTypeTentativeAccept is a specialization of Accept indicating that the acceptance is tentative.
	ActivityTypeTentativeAccept = "TentativeAccept"

	// ActivityTypeTravel indicates that the actor is traveling to target from origin. Travel is an IntransitiveObject whose actor specifies the direct object. If the target or origin are not specified, either can be determined by context.
	ActivityTypeTravel = "Travel"

	// ActivityTypeUndo indicates that the actor is undoing the object. In most cases, the object will be an Activity describing some previously performed action (for instance, a person may have previously "liked" an article but, for whatever reason, might choose to undo that like at some later point in time).
	ActivityTypeUndo = "Undo"

	// ActivityTypeUpdate indicates that the actor has updated the object. Note, however, that this vocabulary does not define a mechanism for describing the actual set of modifications made to object.
	ActivityTypeUpdate = "Update"

	// ActivityTypeView indicates that the actor has viewed the object.
	ActivityTypeView = "View"
)

/**** Actor Types ****/

const (

	// ActorTypeApplication describes a software application.
	ActorTypeApplication = "Application"

	// ActorTypeGroup represents a formal or informal collective of Actors.
	ActorTypeGroup = "Group"

	// ActorTypeOrganization represents an organization.
	ActorTypeOrganization = "Organization"

	// ActorTypePerson represents an individual person.
	ActorTypePerson = "Person"

	// ActorTypeService represents a service of any kind.
	ActorTypeService = "Service"
)

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
