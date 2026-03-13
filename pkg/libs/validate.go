package libs

import "github.com/google/uuid"

func IsUUID(s string) bool {
	if _, err := uuid.Parse(s); err != nil {
		return false
	}
	return true
}
