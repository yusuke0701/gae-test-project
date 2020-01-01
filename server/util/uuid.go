package util

import (
	"github.com/google/uuid"
)

func NewUUID() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
