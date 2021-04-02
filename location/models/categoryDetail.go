package models

import (
	uuid "github.com/satori/go.uuid"
)

//CategoryDetail ...
type CategoryDetail struct {
	Base
	Name           string
	IsActive       bool
	CategoryID     uuid.UUID
	ParentCategory uuid.UUID
}
