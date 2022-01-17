package pub

import (
	"bytes"
	"encoding/json"
)

// toJSON is a shortcut that simplifies the job of marshalling an object by throwing away the
// error that the standard library may throw.  I know, I know, this is usually a bad thing to do,
// so this function should only be called using values that can ALWAYS be marshalled correctly.
func toJSON(value interface{}) []byte {

	result, _ := json.Marshal(value)
	return result
}

func writeToBuffer(buffer bytes.Buffer, items ...interface{}) {

	for _, item := range items {

		switch value := item.(type) {
		case byte:
			buffer.WriteByte(value)

		case rune:
			buffer.WriteRune(value)

		case []byte:
			buffer.Write(value)

		case string:
			buffer.WriteString(value)

		default:
			if json, err := json.Marshal(value); err != nil {
				buffer.Write(json)
			} else {
				buffer.WriteString("null")
			}
		}
	}
}
