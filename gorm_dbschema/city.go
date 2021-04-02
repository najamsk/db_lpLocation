package main

import uuid "github.com/satori/go.uuid"

// City struct
type City struct {
	Base
	CountryID   uuid.UUID
	StateID     uuid.UUID
	Name        string
	DisplayName string
	IsActive    bool
	Locations   []*Location
}
