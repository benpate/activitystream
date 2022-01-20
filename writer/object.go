package writer

import "github.com/whisperverse/pub"

func Article() Object {
	return NewObject().Type(pub.ObjectTypeArticle)
}

func Audio() Object {
	return NewObject().Type(pub.ObjectTypeAudio)
}

func Document() Object {
	return NewObject().Type(pub.ObjectTypeDocument)
}

func Event() Object {
	return NewObject().Type(pub.ObjectTypeEvent)
}

func Image() Object {
	return NewObject().Type(pub.ObjectTypeImage)
}

func Note() Object {
	return NewObject().Type(pub.ObjectTypeNote)
}

func Page() Object {
	return NewObject().Type(pub.ObjectTypePage)
}

func Place() Object {
	return NewObject().Type(pub.ObjectTypePlace)
}

func Profile() Object {
	return NewObject().Type(pub.ObjectTypeProfile)
}

func Relationship() Object {
	return NewObject().Type(pub.ObjectTypeRelationship)
}

func Tombstone() Object {
	return NewObject().Type(pub.ObjectTypeTombstone)
}

func Video() Object {
	return NewObject().Type(pub.ObjectTypeVideo)
}

func Mention() Object {
	return NewObject().Type(pub.ObjectTypeMention)
}
