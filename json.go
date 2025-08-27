package main

import (
	"bytes"
	"encoding/json"
)

type nullable string

func (n *nullable) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("false")) || bytes.Equal(data, []byte("null")) {
		return nil
	}

	var value string

	if err := json.Unmarshal(data, &value); err != nil {
		return nil
	}

	*n = nullable(value)

	return nil
}
