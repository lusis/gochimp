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
	if err != nil {
		tt, err = time.Parse(`"2006-01-02 15:04:05"`, s)
	}
	if err != nil {
		return err
	}
	t.Time = tt
	return nil
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
