package main

import uuid "github.com/satori/go.uuid"

//Staff func
type Staff struct {
	Base
	IsActive      bool
	FirstName     string
	LastName      string
	Email         string
	Mobile        string
	Title         string
	Designation   string
	IsSecondShift bool
	//IsSingleShift bool
}

//LocationStaff func
type LocationStaff struct {
	Base
	UserID     uuid.UUID
	LocationID uuid.UUID
	//IsSingleShift bool
}
