package activitystream

import "encoding/json"

// toJSON is a shortcut that simplifies the job of marshalling an object by throwing away the
// error that the standard library may throw.  I know, I know, this is usually a bad thing to do,
// so this function should only be called using values that can ALWAYS be marshalled correctly.
func toJSON(value interface{}) []byte {

	result, _ := json.Marshal(value)
	return result
}
