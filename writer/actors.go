package writer

import pub "github.com/benpate/activitystream"

func Application(name string, language string) Object {
	return NewObject().
		Type(pub.ActorTypeApplication).
		Name(name, language)
}

func Group(name string, language string) Object {
	return NewObject().
		Type(pub.ActorTypeGroup).
		Name(name, language)
}

func Organization(name string, language string) Object {
	return NewObject().
		Type(pub.ActorTypeOrganization).
		Name(name, language)
}

func Person(name string, language string) Object {
	return NewObject().
		Type(pub.ActorTypePerson).
		Name(name, language)
}

func Service(name string, language string) Object {
	return NewObject().
		Type(pub.ActorTypeService).
		Name(name, language)
}
