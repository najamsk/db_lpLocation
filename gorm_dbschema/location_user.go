package main

import uuid "github.com/satori/go.uuid"

// LocationUser export
type LocationUser struct {
	LocationID uuid.UUID
	UserID     uuid.UUID
}
