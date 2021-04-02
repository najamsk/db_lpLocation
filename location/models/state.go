package models

import (
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	uuid "github.com/satori/go.uuid"
)

// State struct for gorm
type State struct {
	Base
	CountryID   uuid.UUID
	Name        string
	DisplayName string
	IsActive    bool
	Cities      []*City `gorm:"foreignkey:StateID"`
}
