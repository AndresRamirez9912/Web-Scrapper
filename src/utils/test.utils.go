package utils

import "errors"

type errReader int

func (errReader) Read(p []byte) (int, error) {
	return 0, errors.New("Error")
}
