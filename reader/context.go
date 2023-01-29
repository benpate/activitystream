package reader

// Context defines
// https://www.w3.org/TR/json-ld/#the-context
type Context struct {
	Vocabulary  string            `json:"vocabulary"`  // The primary vocabulary represented by the context/document.
	Language    string            `json:"language"`    // The language
	Extensions  map[string]string `json:"extensions"`  // a map of additional namespaces that are included in this context/document.
	NextContext *Context          `json:"nextContext"` // NextContext allows multiple contexts to be defined as a linked list.
}
