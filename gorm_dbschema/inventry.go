package main

import uuid "github.com/satori/go.uuid"

// Inventory struct
type Inventory struct {
	Base
	LocationID uuid.UUID
	Type       string
	count      int64
	Name       string
	TotalCount int64
}
