// Package zuid Zai UUID provides a wrapper around google/uuid to retry creating UUID's if an error occurs.
package zuid

import "github.com/google/uuid"

const MaxRetries = 4

// New creates a random UUID and will retry a maximum of MaxRetries before failing.
func New() (u uuid.UUID, err error) {
	for i := 1; i <= MaxRetries; i++ {
		u, err = uuid.NewRandom()
		if err == nil {
			return u, nil
		}
	}

	return uuid.UUID{}, err
}
