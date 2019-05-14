package activity

const (
	// TypeAccept indicates that the actor accepts the object. The target property can be used in certain circumstances to indicate the context into which the object has been accepted.
	TypeAccept = "Accept"

	// TypeAdd indicates that the actor has added the object to the target. If the target property is not explicitly specified, the target would need to be determined implicitly by context. The origin can be used to identify the context from which the object originated.
	TypeAdd = "Add"

	// TypeAnnounce indicates that the actor is calling the target's attention the object.
	TypeAnnounce = "Announce"

	// TypeArrive is an IntransitiveActivity that indicates that the actor has arrived at the location. The origin can be used to identify the context from which the actor originated. The target typically has no defined meaning.
	TypeArrive = "Arrive"

	// TypeBlock indicates that the actor is blocking the object. Blocking is a stronger form of Ignore. The typical use is to support social systems that allow one user to block activities or content of other users. The target and origin typically have no defined meaning.
	TypeBlock = "Block"

	// TypeCreate indicates that the actor has created the object.
	TypeCreate = "Create"

	// TypeDelete indicates that the actor has deleted the object. If specified, the origin indicates the context from which the object was deleted.
	TypeDelete = "Delete"

	// TypeDislike indicates that the actor dislikes the object.
	TypeDislike = "Dislike"

	// TypeFlag indicates that the actor is "flagging" the object. Flagging is defined in the sense common to many social platforms as reporting content as being inappropriate for any number of reasons.
	TypeFlag = "Flag"

	// TypeFollow indicates that the actor is "following" the object. Following is defined in the sense typically used within Social systems in which the actor is interested in any activity performed by or on the object. The target and origin typically have no defined meaning.
	TypeFollow = "Follow"

	// TypeIgnore indicates that the actor is ignoring the object. The target and origin typically have no defined meaning.
	TypeIgnore = "Ignore"

	// TypeInvite is a specialization of Offer in which the actor is extending an invitation for the object to the target.
	TypeInvite = "Invite"

	// TypeJoin indicates that the actor has joined the object. The target and origin typically have no defined meaning.
	TypeJoin = "Join"

	// TypeLeave indicates that the actor has left the object. The target and origin typically have no meaning.
	TypeLeave = "Leave"

	// TypeLike indicates that the actor likes, recommends or endorses the object. The target and origin typically have no defined meaning.
	TypeLike = "Like"

	// TypeListen indicates that the actor has listened to the object.
	TypeListen = "Listen"

	// TypeMove indicates that the actor has moved object from origin to target. If the origin or target are not specified, either can be determined by context.
	TypeMove = "Move"

	// TypeOffer indicates that the actor is offering the object. If specified, the target indicates the entity to which the object is being offered.
	TypeOffer = "Offer"

	// TypeQuestion represents a question being asked. Question objects are an extension of IntransitiveActivity. That is, the Question object is an Activity, but the direct object is the question itself and therefore it would not contain an object property.
	TypeQuestion = "Question"

	// TypeReject indicates that the actor is rejecting the object. The target and origin typically have no defined meaning.
	TypeReject = "Reject"

	// TypeRead indicates that the actor has read the object.
	TypeRead = "Read"

	// TypeRemove indicates that the actor is removing the object. If specified, the origin indicates the context from which the object is being removed.
	TypeRemove = "Remove"

	// TypeTentativeReject is a specialization of Reject in which the rejection is considered tentative.
	TypeTentativeReject = "TentativeReject"

	// TypeTentativeAccept is a specialization of Accept indicating that the acceptance is tentative.
	TypeTentativeAccept = "TentativeAccept"

	// TypeTravel indicates that the actor is traveling to target from origin. Travel is an IntransitiveObject whose actor specifies the direct object. If the target or origin are not specified, either can be determined by context.
	TypeTravel = "Travel"

	// TypeUndo indicates that the actor is undoing the object. In most cases, the object will be an Activity describing some previously performed action (for instance, a person may have previously "liked" an article but, for whatever reason, might choose to undo that like at some later point in time).
	TypeUndo = "Undo"

	// TypeUpdate indicates that the actor has updated the object. Note, however, that this vocabulary does not define a mechanism for describing the actual set of modifications made to object.
	TypeUpdate = "Update"

	// TypeView indicates that the actor has viewed the object.
	TypeView = "View"
)
