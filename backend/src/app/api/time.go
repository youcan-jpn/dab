package api

import (
	"reflect"
	"time"

	"github.com/youcan-jpn/dab/backend/gen/oapi"
)

var layout = "2006-01-02 15:04:05"

func timeToString(t *time.Time) string {
	str := t.Format(layout)
	return str
}

func objToTime(o *oapi.Date, tz *time.Location) *time.Time {
	if (o == nil) || reflect.ValueOf(o).IsNil() {
		return nil
	}
	y := o.Year
	if *y < 1 || *y > 9999 {
		return nil
	}
	m := time.Month(*o.Month)
	if m < 1 || m > 12 {
		return nil
	}
	d := o.Day
	if *d < 1 || *d > 31 {
		return nil
	}
	t := time.Date(*y, m, *d, 0, 0, 0, 0, tz)
	return &t
}
