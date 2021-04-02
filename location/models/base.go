package models

import (
	"time"

	// _ "github.com/jinzhu/gorm/dialects/postgres"
	uuid "github.com/satori/go.uuid"
)

// Base struct for rest of models
type Base struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;"`
	CreatedBy *uuid.UUID `gorm:"type:uuid;"`
	UpdatedBy *uuid.UUID `gorm:"type:uuid;"`

	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"update_at"`
}
