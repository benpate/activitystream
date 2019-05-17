package object

import "time"

type Event struct {
	Object
	StartTime time.Time
	EndTime   time.Time
}
