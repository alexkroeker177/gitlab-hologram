package util

import "encoding/json"

func ValueFromInterface[T any](i any) (T, error) {
	var v T
	iBts, err := json.Marshal(i)
	if err != nil {
		return v, err
	}
	if err := json.Unmarshal(iBts, &v); err != nil {
		return v, err
	}
	return v, nil
}

func InterfaceFromValue[T any](v T) (any, error) {
	var i any
	iBts, err := json.Marshal(v)
	if err != nil {
		return i, err
	}
	if err := json.Unmarshal(iBts, &i); err != nil {
		return i, err
	}
	return i, nil
}
