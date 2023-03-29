package api

import (
	"time"
)

var layout = "2006-01-02 15:04:05"

func timeToString(t *time.Time) string {
	str := t.Format(layout)
	return str
}
