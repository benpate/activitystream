package writer

import "github.com/benpate/pub"

// Accept
func Accept(actor Object, object Object) Object {
	return NewObject().
		Type(pub.ActivityTypeAccept).
		Actor(actor).
		Object(object)
}

// Add
func Add(actor Object, object Object, target Object) Object {
	return NewObject().
		Type(pub.ActivityTypeAdd).
		Actor(actor).
		Object(object).
		Target(target)
}

// Announce
func Announce(actor Object, object Object, target Object) Object {
	return NewObject().
		Type(pub.ActivityTypeAnnounce).
		Actor(actor).
		Object(object).
		Target(target)
}

// Arrive
func Arrive(actor Object, location Object, origin Object) Object {
	return NewObject().
		Type(pub.ActivityTypeArrive).
		Actor(actor).
		Location(location).
		Origin(origin)
}

// Block
func Block(actor Object, object Object) Object {
	return NewObject().
		Type(pub.ActivityTypeBlock).
		Actor(actor).
		Object(object)
}

// Create
func Create(actor Object, object Object) Object {
	return NewObject().
		Type(pub.ActivityTypeCreate).
		Actor(actor).
		Object(object)
}

// Delete
func Delete(actor Object, object Object, origin Object) Object {
	return NewObject().
		Type(pub.ActivityTypeDelete).
		Actor(actor).
		Object(object).
		Origin(origin)
}

// Dislike
func Dislike(actor Object, object Object) Object {
	return NewObject().
		Type(pub.ActivityTypeDislike).
		Actor(actor).
		Object(object)
}

// Flag
func Flag(actor Object, object Object) Object {
	return NewObject().
		Type(pub.ActivityTypeFlag).
		Actor(actor).
		Object(object)
}

// Follow
func Follow(actor Object, object Object) Object {
	return NewObject().
		Type(pub.ActivityTypeFollow).
		Actor(actor).
		Object(object)
}

// Ignore
func Ignore(actor Object, object Object) Object {
	return NewObject().
		Type(pub.ActivityTypeIgnore).
		Actor(actor).
		Object(object)
}

// Invite
func Invite(actor Object, object Object, target Object) Object {
	return NewObject().
		Type(pub.ActivityTypeInvite).
		Actor(actor).
		Object(object).
		Target(target)
}

// Join
func Join(actor Object, object Object) Object {
	return NewObject().
		Type(pub.ActivityTypeJoin).
		Actor(actor).
		Object(object)
}

// Leave
func Leave(actor Object, object Object) Object {
	return NewObject().
		Type(pub.ActivityTypeLeave).
		Actor(actor).
		Object(object)
}

// Like
func Like(actor Object, object Object) Object {
	return NewObject().
		Type(pub.ActivityTypeLike).
		Actor(actor).
		Object(object)
}

// Listen
func Listen(actor Object, object Object) Object {
	return NewObject().
		Type(pub.ActivityTypeListen).
		Actor(actor).
		Object(object)
}

// Move
func Move(actor Object, object Object, origin Object, target Object) Object {
	return NewObject().
		Type(pub.ActivityTypeMove).
		Actor(actor).
		Object(object).
		Origin(origin).
		Target(target)
}

// Offer
func Offer(actor Object, object Object) Object {
	return NewObject().
		Type(pub.ActivityTypeLike).
		Actor(actor).
		Object(object)
}

// Question
func Question() Object {
	// TODO: this is not complete
	return NewObject().
		Type(pub.ActivityTypeQuestion)
}

// Reject
func Reject(actor Object, object Object) Object {
	return NewObject().
		Type(pub.ActivityTypeReject).
		Actor(actor).
		Object(object)
}

// Read
func Read(actor Object, object Object) Object {
	return NewObject().
		Type(pub.ActivityTypeRead).
		Actor(actor).
		Object(object)
}

// Remove
func Remove(actor Object, object Object, origin Object) Object {
	return NewObject().
		Type(pub.ActivityTypeRemove).
		Actor(actor).
		Object(object).
		Origin(origin)
}

// TentativeAccept
func TentativeAccept(actor Object, object Object) Object {
	return NewObject().
		Type(pub.ActivityTypeTentativeAccept).
		Actor(actor).
		Object(object)
}

// TentativeReject
func TentativeReject(actor Object, object Object) Object {
	return NewObject().
		Type(pub.ActivityTypeTentativeReject).
		Actor(actor).
		Object(object)
}

// Travel
func Travel(actor Object, origin Object, target Object) Object {
	return NewObject().
		Type(pub.ActivityTypeTravel).
		Actor(actor).
		Origin(origin).
		Target(target)
}

// Undo
func Undo(actor Object, object Object) Object {
	return NewObject().
		Type(pub.ActivityTypeUndo).
		Actor(actor).
		Object(object)
}

// Update
func Update(actor Object, object Object) Object {
	return NewObject().
		Type(pub.ActivityTypeUpdate).
		Actor(actor).
		Object(object)
}

// View
func View(actor Object, object Object) Object {
	return NewObject().
		Type(pub.ActivityTypeView).
		Actor(actor).
		Object(object)
}
