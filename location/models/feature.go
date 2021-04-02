package models

import (
	uuid "github.com/satori/go.uuid"
)

//Feature ...
type Feature struct {
	Base
	Name             string
	CategoryDetailID uuid.UUID
	IsActive         bool
	CategoryName     string `gorm:"-" ` // ignore this field while saving
}
