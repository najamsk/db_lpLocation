package models

import (
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	uuid "github.com/satori/go.uuid"
)

// Location expor
type Location struct {
	Base
	Name         string
	DisplayName  string
	IsActive     bool
	CityID       uuid.UUID
	Mobile       string
	GeoLocation  GeoLocation `gorm:"embedded"`
	PhoneNumber  string
	Fax          string
	CountryID    uuid.UUID
	StateID      uuid.UUID
	AddressLine1 string
	AddressLine2 string
	ComapnyID    uuid.UUID
	CountryName  string `gorm:"-"`  // ignore this field while saving
	StateName    string `gorm:"-"`  // ignore this field while saving
	CityName     string `gorm:"-" `  // ignore this field while saving
	//Users       []*User     `gorm:"many2many:location_users;"`
}
