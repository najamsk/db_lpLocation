package main

import uuid "github.com/satori/go.uuid"

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
	//Users       []*User     `gorm:"many2many:location_users;"`
}

//TODO: location shift enum {first, extended, full}
