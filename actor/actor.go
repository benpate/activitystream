package actor

import "github.com/benpate/activitystream"

// Application returns an `Actor` object of type `Application`
func Application(name string) activitystream.Object {
	return Object{
		Context: activitystream.ContextDefault(),
		Type:    TypeApplication,
		Name:    name,
	}
}

// Group returns an `Actor` object of type `Group`
func Group(name string) activitystream.Object {
	return Object{
		Context: activitystream.ContextDefault(),
		Type:    TypeGroup,
		Name:    name,
	}
}

// Organization returns an `Actor` object of type `Organization`
func Organization(name string) activitystream.Object {
	return Object{
		Context: activitystream.ContextDefault(),
		Type:    TypeOrganization,
		Name:    name,
	}
}

// Person represents an individual person.
func Person(name string) activitystream.Object {
	return Object{
		Context: activitystream.ContextDefault(),
		Type:    TypePerson,
		Name:    name,
	}
}

// Service represents a service of any kind
func Service(name string) activitystream.Object {
	return Object{
		Context: activitystream.ContextDefault(),
		Type:    TypeService,
		Name:    name,
	}
}
