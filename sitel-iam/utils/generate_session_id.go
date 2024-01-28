package utils

import (
	"github.com/google/uuid"
)

func GenerateSessionId() string {
	return uuid.New().String()
}
