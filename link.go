package activitystream

func Link(href string) *Object {
	return &Object{
		Type: []string{"link"},
		Href: href,
	}
}
