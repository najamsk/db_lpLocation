package main

import uuid "github.com/satori/go.uuid"

// LocationProductType export:w
type LocationProductType struct {
	Base
	LocationID       uuid.UUID
	CategoryDetailID uuid.UUID // meeting/conf/private room not actual product
	IsActive         bool
}
