package object

import "time"

type Tombstone struct {
	Object

	FormerType string
	Deleted    time.Time
}
