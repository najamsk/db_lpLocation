package main

// Country struct
type Country struct {
	Base
	Name        string
	DisplayName string
	IsActive    bool
	States      []*State `gorm:"foreignkey:CountryID"`
	//Cities      []*City  `gorm:"foreignkey:CountryID"`
}
