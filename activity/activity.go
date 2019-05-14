package activity

import "github.com/benpate/activitystream"

	// Accept indicates that the actor accepts the object. The target property can be used in certain circumstances to indicate the context into which the object has been accepted.
	func Accept(actor *activitystream.Object, object *activitystream.Object, target *activitystream.Object) activitystream.Object {
		return activitystream.Object{
			Context: activityStream.ContextDefault(),
			Type: TypeAccept,
		}

	// Add indicates that the actor has added the object to the target. If the target property is not explicitly specified, the target would need to be determined implicitly by context. The origin can be used to identify the context from which the object originated.
	func Add(actor *activitystream.Object, object *activitystream.Object, target *activitystream.Object, oritin *activitystream.Object) activitystream.Object {
		return activitystream.Object{
			Context: activityStream.ContextDefault(),
			Type: TypeAdd,
		}

	// Announce indicates that the actor is calling the target's attention the object.
	func Announce(actor *activitystream.Object, object *activitystream.Object, target *activitystream.Object) activitystream.Object {
		return activitystream.Object{
			Context: activityStream.ContextDefault(),
			Type: TypeAnnounce,
		}

	// Arrive is an IntransitiveActivity that indicates that the actor has arrived at the location. The origin can be used to identify the context from which the actor originated. The target typically has no defined meaning.
	func Arrive(actor *activitystream.Object, object *activitystream.Object, location *activitystream.Object) activitystream.Object {
		return activitystream.Object{
			Context: activityStream.ContextDefault(),
			Type: TypeArrive,
		}

	// Block indicates that the actor is blocking the object. Blocking is a stronger form of Ignore. The typical use is to support social systems that allow one user to block activities or content of other users. The target and origin typically have no defined meaning.
	func Block(actor *activitystream.Object, object *activitystream.Object, target *activitystream.Object) activitystream.Object {
		return activitystream.Object{
			Context: activityStream.ContextDefault(),
			Type: TypeBlock,
		}

	// Create indicates that the actor has created the object.
	func Create(actor *activitystream.Object, object *activitystream.Object) activitystream.Object {
		return activitystream.Object{
			Context: activityStream.ContextDefault(),
			Type: TypeCreate,
		}

	// Delete indicates that the actor has deleted the object. If specified, the origin indicates the context from which the object was deleted.
	func Delete(actor *activitystream.Object, object *activitystream.Object, origin *activitystream.Object) activitystream.Object {
		return activitystream.Object{
			Context: activityStream.ContextDefault(),
			Type: TypeDelete,
		}

	// Dislike indicates that the actor dislikes the object.
	func Dislike(actor *activitystream.Object, object *activitystream.Object, target *activitystream.Object) activitystream.Object {
		return activitystream.Object{
			Context: activityStream.ContextDefault(),
			Type: TypeDislike,
		}

	// Flag indicates that the actor is "flagging" the object. Flagging is defined in the sense common to many social platforms as reporting content as being inappropriate for any number of reasons.
	func Flag(actor *activitystream.Object, object *activitystream.Object, target *activitystream.Object) activitystream.Object {
		return activitystream.Object{
			Context: activityStream.ContextDefault(),
			Type: TypeFlag,
		}

	// Follow indicates that the actor is "following" the object. Following is defined in the sense typically used within Social systems in which the actor is interested in any activity performed by or on the object. The target and origin typically have no defined meaning.
	func Follow(actor *activitystream.Object, object *activitystream.Object) activitystream.Object {
		return activitystream.Object{
			Context: activityStream.ContextDefault(),
			Type: TypeFollow,
		}

	// Ignore indicates that the actor is ignoring the object. The target and origin typically have no defined meaning.
	func Ignore(actor *activitystream.Object, object *activitystream.Object) activitystream.Object {
		return activitystream.Object{
			Context: activityStream.ContextDefault(),
			Type: TypeIgnore,
		}

	// Invite is a specialization of Offer in which the actor is extending an invitation for the object to the target.
	func Invite(actor *activitystream.Object, object *activitystream.Object, target *activitystream.Object) activitystream.Object {
		return activitystream.Object{
			Context: activityStream.ContextDefault(),
			Type: TypeInvite,
		}

	// Join indicates that the actor has joined the object. The target and origin typically have no defined meaning.
	func Join(actor *activitystream.Object, object *activitystream.Object) activitystream.Object {
		return activitystream.Object{
			Context: activityStream.ContextDefault(),
			Type: TypeJoin,
		}

	// Leave indicates that the actor has left the object. The target and origin typically have no meaning.
	func Leave(actor *activitystream.Object, object *activitystream.Object) activitystream.Object {
		return activitystream.Object{
			Context: activityStream.ContextDefault(),
			Type: TypeLeave,
		}

	// Like indicates that the actor likes, recommends or endorses the object. The target and origin typically have no defined meaning.
	func Like(actor *activitystream.Object, object *activitystream.Object) activitystream.Object {
		return activitystream.Object{
			Context: activityStream.ContextDefault(),
			Type: TypeLike,
		}

	// Listen indicates that the actor has listened to the object.
	func Listen(actor *activitystream.Object, object *activitystream.Object) activitystream.Object {
		return activitystream.Object{
			Context: activityStream.ContextDefault(),
			Type: TypeListen,
		}

	// Move indicates that the actor has moved object from origin to target. If the origin or target are not specified, either can be determined by context.
	func Move(actor *activitystream.Object, object *activitystream.Object, target *activitystream.Object) activitystream.Object {
		return activitystream.Object{
			Context: activityStream.ContextDefault(),
			Type: TypeMove,
		}

	// Offer indicates that the actor is offering the object. If specified, the target indicates the entity to which the object is being offered.
	func Offer(actor *activitystream.Object, object *activitystream.Object, target *activitystream.Object) activitystream.Object {
		return activitystream.Object{
			Context: activityStream.ContextDefault(),
			Type: TypeOffer,
		}

	// Question represents a question being asked. Question objects are an extension of IntransitiveActivity. That is, the Question object is an Activity, but the direct object is the question itself and therefore it would not contain an object property.
	func Question(actor *activitystream.Object, object *activitystream.Object, target *activitystream.Object) activitystream.Object {
		return activitystream.Object{
			Context: activityStream.ContextDefault(),
			Type: TypeQuestion,
		}

	// Reject indicates that the actor is rejecting the object. The target and origin typically have no defined meaning.
	func Reject(actor *activitystream.Object, object *activitystream.Object) activitystream.Object {
		return activitystream.Object{
			Context: activityStream.ContextDefault(),
			Type: TypeReject,
		}

	// Read indicates that the actor has read the object.
	func Read(actor *activitystream.Object, object *activitystream.Object, target *activitystream.Object) activitystream.Object {
		return activitystream.Object{
			Context: activityStream.ContextDefault(),
			Type: TypeRead,
		}

	// Remove indicates that the actor is removing the object. If specified, the origin indicates the context from which the object is being removed.
	func Remove(actor *activitystream.Object, object *activitystream.Object, target *activitystream.Object) activitystream.Object {
		return activitystream.Object{
			Context: activityStream.ContextDefault(),
			Type: TypeRemove,
		}

	// TentativeReject is a specialization of Reject in which the rejection is considered tentative.
	func TentativeReject(actor *activitystream.Object, object *activitystream.Object) activitystream.Object {
		return activitystream.Object{
			Context: activityStream.ContextDefault(),
			Type: TypeTentativeReject,
		}

	// TentativeAccept is a specialization of Accept indicating that the acceptance is tentative.
	func TentativeAccept(actor *activitystream.Object, object *activitystream.Object) activitystream.Object {
		return activitystream.Object{
			Context: activityStream.ContextDefault(),
			Type: TypeTentativeAccept,
		}

	// Travel indicates that the actor is traveling to target from origin. Travel is an IntransitiveObject whose actor specifies the direct object. If the target or origin are not specified, either can be determined by context.
	func Travel(actor *activitystream.Object, object *activitystream.Object, target *activitystream.Object) activitystream.Object {
		return activitystream.Object{
			Context: activityStream.ContextDefault(),
			Type: TypeTravel,
		}

	// Undo indicates that the actor is undoing the object. In most cases, the object will be an Activity describing some previously performed action (for instance, a person may have previously "liked" an article but, for whatever reason, might choose to undo that like at some later point in time).
	func Undo(actor *activitystream.Object, object *activitystream.Object) activitystream.Object {
		return activitystream.Object{
			Context: activityStream.ContextDefault(),
			Type: TypeUndo,
		}

	// Update indicates that the actor has updated the object. Note, however, that this vocabulary does not define a mechanism for describing the actual set of modifications made to object.
	func Update(actor *activitystream.Object, object *activitystream.Object) activitystream.Object {
		return activitystream.Object{
			Context: activityStream.ContextDefault(),
			Type: TypeUpdate,
		}

	// View indicates that the actor has viewed the object.
	func View(actor *activitystream.Object, object *activitystream.Object) activitystream.Object {
		return activitystream.Object{
			Context: activityStream.ContextDefault(),
			Type: TypeView,
		}
)