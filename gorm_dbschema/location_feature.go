package main

import uuid "github.com/satori/go.uuid"

// LocationFeature export
type LocationFeature struct {
	Base
	FeatureID  uuid.UUID
	IsActive   bool
	Price      float64
	LocationID uuid.UUID
	IsAddOn    bool // if true price will kick in
}
