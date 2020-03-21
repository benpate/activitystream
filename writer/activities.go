package writer

// Accept
func Accept(actor *Object, object *Object) *Object {}

// Add
func Add(actor *Object, *object *Object, target *Object) *Object {}

// Announce
func Announce(actor *Object, object *Object, target *Object) *Object {}

// Arrive
func Arrive(actor *Object, location *Object, origin *Object) *Object {}

// Block
func Block(actor *Object, object *Object) *Object {}

// Create
func Create(actor *Object, *object *Object) {}

// Delete
func Delete(actor *Object, *object *Object, origin *Object) {}

// Dislike
func Dislike(actor *Object, object *Object) *Object {}

// Flag
func Flag(actor *Object, object *Object) *Object {}

// Follow
func Follow(actor *Object, object *Object) *Object {}

// Ignore
func Ignore(actor *Object, object *Object) *Object {}

// Invite
func Invite(actor *Object, object *Object, target *Object) *Object {}

// Join
func Join(actor *Object, object *Object) *Object {}

// Leave
func Leave(actor *Object, object *Object) *Object {}

// Like
func Like(actor *Object, object *Object) *Object {}

// Listen
func Listen(actor *Object, object *Object) *Object {}

// Move
func Move(actor *Object, object *Object, origin *Object, target *Object) *Object {}

// Offer
func Offer(actor *Object, object *Object) *Object {}

// Question
func Question() *Object {}

// Reject
func Reject(actor *Object, object *Object) *Object {}

// Read
func Read(actor *Object, object *Object) *Object {}

// Remove
func Remove(actor *Object, object *Object, origin *Object) *Object {}

// TentativeAccept
func TentativeAccept(actor *Object, *object *Object) *Object {}

// TentativeReject
func TentativeReject(actor *Object, object *Object) *Object {}

// Travel
func Travel(actor *Object, origin *Object, target *Object) *Object {}

// Undo
func Undo(actor *Object, object *Object) *Object {}

// Update
func Update(actor *Object, object *Object) *Object {}

// View
func View(actor *Object, object *Object) *Object {}
