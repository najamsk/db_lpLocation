package models

import (
// _ "github.com/jinzhu/gorm/dialects/postgres"
)

// GeoLocation struct
type GeoLocation struct {
	GeoLocationLat  float64
	GeoLocationLong float64
	Radius          float64
}
