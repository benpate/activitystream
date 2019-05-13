package model

type Collection struct {
	Object
	TotalItems int
	Current    string
	First      string
	Last       string
	Items      []Object
}
