package main

import (
	uuid "github.com/satori/go.uuid"
)

// Feature type
//type Feature struct {
//	Base
//	EntityID          uuid.UUID //category id
//	IsActive          bool
//	IsPackageFeature  bool
//	IsProductFeature  bool
//	IsLocationFeature bool
//}

// Feature type
type Feature struct {
	Base
	Name             string    //lunch or dinner two seprate entries
	CategoryDetailID uuid.UUID //Drinks
	IsActive         bool
}

// Category type
type Category struct {
	Base
	Name     string
	IsActive bool
}

// CategoryDetail type
type CategoryDetail struct {
	Base
	Name           string
	IsActive       bool
	CategoryID     uuid.UUID
	ParentCategory uuid.UUID
}
