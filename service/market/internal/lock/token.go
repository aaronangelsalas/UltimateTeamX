package lock

import "github.com/google/uuid"

func newToken() string {
	return uuid.NewString()
}
