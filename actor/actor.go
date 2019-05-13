package actor

// Application returns an `Actor` object of type `Application`
func Application(name string) Object {
	return Object{
		Context: ContextActivityStreams(),
		Type:    ActorTypeApplication,
		Name:    name,
	}
}

// Group returns an `Actor` object of type `Group`
func Group(name string) Object {
	return Object{
		Context: ContextActivityStreams(),
		Type:    ActorTypeGroup,
		Name:    name,
	}
}

// Organization returns an `Actor` object of type `Organization`
func Organization(name string) Object {
	return Object{
		Context: ContextActivityStreams(),
		Type:    ActorTypeOrganization,
		Name:    name,
	}
}

func Person(name string) Object {
	return Object{
		Context: ContextActivityStreams(),
		Type:    ActorTypePerson,
		Name:    name,
	}
}

func Service(name string) Object {
	return Object{
		Context: ContextActivityStreams(),
		Type:    ActorTypeService,
		Name:    name,
	}
}
