package api

import (
	"bytes"
	"encoding/json"
)

func testStrictDecode(b []byte, t interface{}) error {
	decoder := json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(t); err != nil {
		return err
	}
	return nil
}
