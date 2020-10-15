package api

import (
	"time"
)

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(data []byte) error {
	s := string(data)
	tt, err := time.Parse(`"`+APITimeFormat+`"`, s)
	t.Time = tt
	return err
}

//APITimeFormat is the format string for time.Format
const APITimeFormat = time.RFC3339

// nolint: deadcode, varcheck, unused
func apiTime(t interface{}) interface{} {
	switch ti := t.(type) {
	case time.Time:
		return ti.Format(APITimeFormat)
	case string:
		return ti
	}
	return t
}
