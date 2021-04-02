package main

import uuid "github.com/satori/go.uuid"

// LocationGallery export
type LocationGallery struct {
	Base
	LocationID uuid.UUID
	MediaURL   string
	IsActive   bool
}
