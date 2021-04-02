package main

import (
	"fmt"
	"time"

	//"github.com/gogo/protobuf/proto"

	"errors"

	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	gservice "github.com/launchpad7/lp7_micro_location/grpc"
	"github.com/launchpad7/lp7_micro_location/location"
	locationmodel "github.com/launchpad7/lp7_micro_location/location/models"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// LocationGRPCServer struct
type LocationGRPCServer struct {
	locationService location.Service
}

// NewLocationGRPCServer return the grpc servcie enabled server
func NewLocationGRPCServer(locationService location.Service) *LocationGRPCServer {
	fmt.Println("new location grpc server invoked")
	return &LocationGRPCServer{locationService: locationService}
}

// AllCountries func
func (r LocationGRPCServer) AllCountries(ctx context.Context, request *gservice.FilterLocation) (*gservice.CountryList, error) {
	//a := request.GetHello()

	//getting header to print requestid
	header, _ := metadata.FromIncomingContext(ctx)
	headerRID := header["reqid"]
	fmt.Printf("ReqIDHeader = %+v \n", headerRID)

	//getting request entity id from responsebase
	rCountryID := request.GetCountryid()

	countries, errCountries := r.locationService.GetCountries(rCountryID)

	if errCountries != nil {
		fmt.Printf("error returning countries from locationService.GetCountryList: %+v \n", errCountries)
	}
	//fmt.Printf("countries return from hex-location is :%+v \n", countries)
	gcountries := []*gservice.Country{}

	if countries == nil {
		fmt.Printf("countries return from hlocation is nil \n")
	}
	for _, c := range *countries {
		fmt.Printf("country = %+v \n", c)
		fmt.Printf("country createdat unix = %+v \n", c.CreatedAt.Unix())

		var createdatStamp timestamp.Timestamp
		createdatStamp.Seconds = c.CreatedAt.Unix()

		// createdatStamp.set_seconds(c.CreatedAt.Unix())

		gbase := &gservice.ResponseBase{Id: c.ID.String(), Createdatunix: &createdatStamp, Createdat: c.CreatedAt.String(), Updatedat: c.UpdatedAt.String()}
		gcountries = append(gcountries, &gservice.Country{Countryname: c.Name, Isactive: c.IsActive, B: gbase})

	}

	return &gservice.CountryList{Countries: gcountries}, nil
}

// AllStates func
func (r LocationGRPCServer) AllStates(ctx context.Context, request *gservice.FilterLocation) (*gservice.StateList, error) {
	//getting header to print requestid
	header, _ := metadata.FromIncomingContext(ctx)
	//devHeaderValue := header["dev"]
	//networkHeaderValue := header["network"]
	headerRID := header["reqid"]
	fmt.Printf("ReqIDHeader = %+v \n", headerRID)

	//getting request entity id from responsebase
	rCountryID := request.GetCountryid()
	rStateID := request.GetStateid()
	//rCityID := request.GetCityid()

	fmt.Printf("rCountryID = %v \n", rCountryID)
	fmt.Printf("rStateID = %v \n", rStateID)
	//fmt.Printf("rCityID = %v \n", rCityID)

	states, errStates := r.locationService.GetStates(rStateID, rCountryID)
	if errStates != nil {
		fmt.Printf("error returning countries from locationService.GetCountryList: %+v \n", errStates)
	}
	//fmt.Printf("countries return from hex-location is :%+v \n", countries)
	gstates := []*gservice.State{}

	if states == nil {
		fmt.Printf("states return from hlocation is nil \n")
	}
	for _, c := range *states {
		fmt.Printf("country = %+v \n", c)
		var createdatStamp timestamp.Timestamp
		createdatStamp.Seconds = c.CreatedAt.Unix()

		// createdatStamp.set_seconds(c.CreatedAt.Unix())

		gbase := &gservice.ResponseBase{
			Id:            c.ID.String(),
			Createdatunix: &createdatStamp,
			Createdat:     c.CreatedAt.String(),
			Updatedat:     c.UpdatedAt.String()}
		//countries = append(gcountries, &gservice.Country{Countryname: c.Name, Isactive: c.IsActive, B: gbase})
		gstates = append(gstates, &gservice.State{
			Statename: c.Name,
			Countryid: c.CountryID.String(),
			Isactive:  c.IsActive,
			B:         gbase})

	}

	//lookup db for passed key if no results return empty list
	//gstates := []*gservice.State{
	//	&gservice.State{Statename: "sindh", Countryid: "092", Countryname: "Pakistan", Isactive: false},
	//	&gservice.State{Statename: "punjab", Countryid: "092", Countryname: "Pakistan", Isactive: true},
	//	&gservice.State{Statename: "islamabad", Countryid: "092", Countryname: "Pakistan", Isactive: true},
	//	&gservice.State{Statename: "balochistan", Countryid: "092", Countryname: "Pakistan", Isactive: false},
	//	&gservice.State{Statename: "kpk", Countryid: "092", Countryname: "Pakistan", Isactive: false},
	//}

	return &gservice.StateList{States: gstates}, nil
}

// AllCities func
func (r LocationGRPCServer) AllCities(ctx context.Context, request *gservice.FilterLocation) (*gservice.CityList, error) {

	//getting header to print requestid
	header, _ := metadata.FromIncomingContext(ctx)
	headerRID := header["reqid"]
	fmt.Printf("ReqIDHeader = %+v \n", headerRID)

	//getting request entity id from responsebase
	rCountryID := request.GetCountryid()
	rStateID := request.GetStateid()
	rCityID := request.GetCityid()

	fmt.Printf("rCountryID = %v \n", rCountryID)
	fmt.Printf("rStateID = %v \n", rStateID)
	fmt.Printf("rCityID = %v \n", rCityID)

	cities, errCities := r.locationService.GetCities(rCityID, rStateID, rCountryID)
	if errCities != nil {
		fmt.Printf("error returning cities from locationService.GetCountryList: %+v \n", errCities)
	}
	//fmt.Printf("countries return from hex-location is :%+v \n", countries)
	gcities := []*gservice.City{}

	if cities == nil {
		fmt.Printf("cities return from hlocation is nil \n")
	}

	for _, c := range *cities {
		fmt.Printf("city = %+v \n", c)

		var createdatStamp timestamp.Timestamp
		createdatStamp.Seconds = c.CreatedAt.Unix()

		// createdatStamp.set_seconds(c.CreatedAt.Unix())

		gbase := &gservice.ResponseBase{
			Id:            c.ID.String(),
			Createdatunix: &createdatStamp,
			Createdat:     c.CreatedAt.String(),
			Updatedat:     c.UpdatedAt.String()}

		gcities = append(gcities, &gservice.City{
			Cityname:  c.Name,
			Countryid: c.CountryID.String(),
			Stateid:   c.StateID.String(),
			Isactive:  c.IsActive,
			B:         gbase})

	}

	//cities := []*gservice.City{
	//	&gservice.City{Cityname: "karachi"},
	//	&gservice.City{Cityname: "Lahore"},
	//	&gservice.City{Cityname: "Multan"},
	//	&gservice.City{Cityname: "Islamabad"},
	//	&gservice.City{Cityname: "Rawalpindi"},
	//}

	reslog := fmt.Sprintf("cities result:  %v", cities)
	fmt.Println("cities = ", reslog)
	return &gservice.CityList{Cities: gcities}, nil
}

// AllLocations func
func (r LocationGRPCServer) AllLocations(ctx context.Context, request *gservice.FilterLocation) (*gservice.LocationList, error) {
	//getting header to print requestid
	header, _ := metadata.FromIncomingContext(ctx)
	//devHeaderValue := header["dev"]
	//networkHeaderValue := header["network"]
	headerRID := header["reqid"]
	fmt.Printf("ReqIDHeader = %+v \n", headerRID)

	//getting request entity id from responsebase
	rCountryID := request.GetCountryid()
	rStateID := request.GetStateid()
	rCityID := request.GetCityid()
	rLocationID := request.GetLocationid()

	fmt.Printf("rCountryID = %v \n", rCountryID)
	fmt.Printf("rStateID = %v \n", rStateID)
	fmt.Printf("rCityID = %v \n", rCityID)
	fmt.Printf("rLocationID = %v \n", rLocationID)

	//lookup db for passed key if no results return empty list
	locations, errLocations := r.locationService.GetLocations(rLocationID, rCityID, rStateID, rCountryID)

	if errLocations != nil {
		fmt.Printf("error returning locations from locationService.GetCountryList: %+v \n", errLocations)
	}
	//fmt.Printf("countries return from hex-location is :%+v \n", countries)
	glocations := []*gservice.Location{}

	if locations == nil {
		fmt.Printf("locations return from hlocation is nil \n")
	}

	gFeature := []*gservice.Feature{}
	gallFeature := []*gservice.Feature{}
	if len(rLocationID) > 0 {
		locationFeatures2, _ := r.locationService.GetLocationFeatures(rLocationID)
		for _, f := range *locationFeatures2 {

			gbase := &gservice.ResponseBase{
				Id:        f.ID.String(),
				Createdat: f.CreatedAt.String(),
				Updatedat: f.UpdatedAt.String()}

			gFeature = append(gFeature, &gservice.Feature{
				Name: f.Name,
				B:    gbase,
			})

		}

		allFeatures, _ := r.locationService.GetFeatures("")
		for _, f := range *allFeatures {

			gbase := &gservice.ResponseBase{
				Id:        f.ID.String(),
				Createdat: f.CreatedAt.String(),
				Updatedat: f.UpdatedAt.String()}

			gallFeature = append(gallFeature, &gservice.Feature{
				Name: f.Name,
				B:    gbase,
			})

		}
	}

	fmt.Println("allFeatures:", gallFeature)
	fmt.Println("gFeature:", gFeature)
	for _, c := range *locations {
		fmt.Printf("city = %+v \n", c)

		var createdatStamp timestamp.Timestamp
		if c.CreatedAt != nil {
			createdatStamp.Seconds = c.CreatedAt.Unix()
		}

		var updatedatStamp timestamp.Timestamp
		if c.UpdatedAt != nil {
			updatedatStamp.Seconds = c.UpdatedAt.Unix()
		}

		// createdatStamp.set_seconds(c.CreatedAt.Unix())

		gbase := &gservice.ResponseBase{
			Id:            c.ID.String(),
			Createdatunix: &createdatStamp,
			Createdat:     c.CreatedAt.String(),
			Updatedat:     c.UpdatedAt.String(),
			Updatedatunix: &updatedatStamp,
		}

		if c.CreatedBy != nil {
			gbase.Createdby = c.CreatedBy.String()
		}
		if c.UpdatedBy != nil {
			gbase.Updatedby = c.UpdatedBy.String()
		}

		glocations = append(glocations, &gservice.Location{
			Name:         c.Name,
			Displayname:  c.DisplayName,
			Isactive:     c.IsActive,
			Addressline1: c.AddressLine1,
			Addressline2: c.AddressLine2,
			Fax:          c.Fax,
			Phonenumber:  c.PhoneNumber,
			Mobile:       c.Mobile,
			Locality: &gservice.Locality{
				Countryid:   c.CountryID.String(),
				Stateid:     c.StateID.String(),
				Cityid:      c.CityID.String(),
				Countryname: c.CountryName,
				Statename:   c.StateName,
				Cityname:    c.CityName,
			},
			B: gbase,
			Gloc: &gservice.GeoLocation{
				Lat:    c.GeoLocation.GeoLocationLat,
				Lon:    c.GeoLocation.GeoLocationLong,
				Radius: c.GeoLocation.Radius,
			},
			Locationfeatures: &gservice.FeatureList{Features: gFeature},
			Allfeatures:      &gservice.FeatureList{Features: gallFeature},
		})

	}
	fmt.Println("glocations:", glocations)
	return &gservice.LocationList{Locations: glocations}, nil
}

// AllCategories func
func (r LocationGRPCServer) AllCategories(ctx context.Context, request *gservice.FilterCategory) (*gservice.CategoryList, error) {

	// gcategories := []*gservice.Category{
	// 	&gservice.Category{Name: "Drinks", Isactive: true},
	// 	&gservice.Category{Name: "Meals", Isactive: true},
	// }
	// return &gservice.CategoryList{Categories: gcategories}, nil

	//getting header to print requestid
	header, _ := metadata.FromIncomingContext(ctx)
	headerRID := header["reqid"]
	fmt.Printf("ReqIDHeader = %+v \n", headerRID)

	//getting request entity id from responsebase
	rCategoryID := request.GetCategoryid()

	fmt.Printf("rCategoryDetailID = %v \n", rCategoryID)

	category, errCat := r.locationService.GetCategory(rCategoryID)
	if errCat != nil {
		fmt.Printf("error returning category detail from locationService.GetCategory: %+v \n", errCat)
	}
	//fmt.Printf("countries return from hex-location is :%+v \n", countries)
	gcategory := []*gservice.Category{}

	if category == nil {
		fmt.Printf("category return from hlocation is nil \n")
	}

	for _, c := range *category {
		fmt.Printf("city = %+v \n", c)

		var createdatStamp timestamp.Timestamp
		createdatStamp.Seconds = c.CreatedAt.Unix()

		// createdatStamp.set_seconds(c.CreatedAt.Unix())

		gbase := &gservice.ResponseBase{
			Id:            c.ID.String(),
			Createdatunix: &createdatStamp,
			Createdat:     c.CreatedAt.String(),
			Updatedat:     c.UpdatedAt.String()}

		gcategory = append(gcategory, &gservice.Category{
			Name:     c.Name,
			Isactive: c.IsActive,
			B:        gbase})

	}

	reslog := fmt.Sprintf("category result:  %v", category)
	fmt.Println("cities = ", reslog)
	return &gservice.CategoryList{Categories: gcategory}, nil
}

// AllCategoryDetails func
// func (r LocationGRPCServer) AllCategoryDetails(ctx context.Context, request *gservice.FilterCategoryDetail) (*gservice.CategoryDetailList, error) {

// 	gcategorydetails := []*gservice.CategoryDetail{
// 		&gservice.CategoryDetail{Name: "Lassi", Isactive: true, Categoryid: "1", Categoryname: "Drinks"},
// 		&gservice.CategoryDetail{Name: "Tea", Isactive: true, Categoryid: "1", Categoryname: "Drinks"},
// 		&gservice.CategoryDetail{Name: "Lunch", Isactive: true, Categoryid: "2", Categoryname: "Meals"},
// 		&gservice.CategoryDetail{Name: "Dinner", Isactive: true, Categoryid: "2", Categoryname: "Meals"},
// 	}

// 	return &gservice.CategoryDetailList{Categorydetails: gcategorydetails}, nil
// }

// AllFeatures func
func (r LocationGRPCServer) AllFeatures(ctx context.Context, request *gservice.FilterFeature) (*gservice.FeatureList, error) {

	//static response for now -
	// gfeatures := []*gservice.Feature{
	// 	&gservice.Feature{Name: "namkeen lassi", Isactive: true, Categorydetailid: "1", Categorydetailname: "Lassi"},
	// 	&gservice.Feature{Name: "Meals", Isactive: true, Categorydetailid: "1", Categorydetailname: "Lassi"},
	// 	&gservice.Feature{Name: "Reserved Car Parking", Isactive: true, Categorydetailid: "3", Categorydetailname: "Car Parking"},
	// 	&gservice.Feature{Name: "First Come, Parking", Isactive: false, Categorydetailid: "4", Categorydetailname: "Bike Parking"},
	// }
	// return &gservice.FeatureList{Features: gfeatures}, nil
	rID := request.GetFeatureid()
	gallFeature := []*gservice.Feature{}
	allFeatures, _ := r.locationService.GetFeatures(rID)
	for _, f := range *allFeatures {

		var createdatStamp timestamp.Timestamp
		if f.CreatedAt != nil {
			createdatStamp.Seconds = f.CreatedAt.Unix()
		}

		var updatedatStamp timestamp.Timestamp
		if f.UpdatedAt != nil {
			updatedatStamp.Seconds = f.UpdatedAt.Unix()
		}
		gbase := &gservice.ResponseBase{
			Id:            f.ID.String(),
			Createdatunix: &createdatStamp,
			Createdat:     f.CreatedAt.String(),
			Updatedat:     f.UpdatedAt.String(),
			Updatedatunix: &updatedatStamp,
		}

		if f.CreatedBy != nil {
			gbase.Createdby = f.CreatedBy.String()
		}
		if f.UpdatedBy != nil {
			gbase.Updatedby = f.UpdatedBy.String()
		}

		gallFeature = append(gallFeature, &gservice.Feature{
			Name:               f.Name,
			Categorydetailid:   f.CategoryDetailID.String(),
			Categorydetailname: f.CategoryName,
			Isactive:           f.IsActive,
			B:                  gbase,
		})

	}
	return &gservice.FeatureList{Features: gallFeature}, nil
}

// AllStaff func
func (r LocationGRPCServer) AllStaff(ctx context.Context, request *gservice.FilterStaff) (*gservice.StaffList, error) {

	//static response for now
	gstaff := []*gservice.StaffMember{
		&gservice.StaffMember{
			Userid:        "udb1",
			Locationid:    "loc1",
			Isactive:      true,
			Firstname:     "zohaib",
			Lastname:      "Zehreela",
			Email:         "zz@jsp.com",
			Mobile:        "0300-368888",
			Title:         "Mr",
			Designation:   "Media Guru",
			Issecondshift: false,
		},
		&gservice.StaffMember{
			Userid:        "udb2",
			Locationid:    "loc2",
			Isactive:      true,
			Firstname:     "shakeela",
			Lastname:      "sadaqat",
			Email:         "shakeela@jsp.com",
			Mobile:        "0300-98989",
			Title:         "Miss",
			Designation:   "Fine Officer",
			Issecondshift: false,
		},
	}
	return &gservice.StaffList{Members: gstaff}, nil
}

// AllLocationFeatures func
func (r LocationGRPCServer) AllLocationFeatures(ctx context.Context, request *gservice.FilterFeature) (*gservice.FeatureList, error) {
	return nil, errors.New("not implemented yet")
}

// SaveLocation func
func (r LocationGRPCServer) SaveLocation(ctx context.Context,
	request *gservice.Location) (*gservice.Location, error) {
	//getting header to print requestid
	header, _ := metadata.FromIncomingContext(ctx)
	headerRID := header["reqid"]
	fmt.Printf("ReqIDHeader = %+v \n", headerRID)

	//Request Parsing
	var createddate time.Time
	var rLocationID uuid.UUID
	var rUserID uuid.UUID
	var rCreatedUserID uuid.UUID
	rName := request.GetName()
	if len(rName) == 0 {
		return nil, errors.New("name is required")
	}
	rDisplayname := request.GetDisplayname()
	if len(rDisplayname) == 0 {
		return nil, errors.New("displayname is required")
	}
	rIsactive := request.GetIsactive()
	rGloc := request.GetGloc()
	rBase := request.GetB()
	rLocality := request.GetLocality()
	rAddressline1 := request.GetAddressline1()
	featureids := request.GetLocationfeatureids()

	if len(rAddressline1) == 0 {
		return nil, errors.New("AddressLine1 is required")
	}
	rAddressline2 := request.GetAddressline2()
	if len(rAddressline2) == 0 {
		return nil, errors.New("AddressLine2 is required")
	}
	rCityid := rLocality.GetCityid()
	if len(rCityid) == 0 {
		return nil, errors.New("Cityid is required")
	}
	rStateid := rLocality.GetStateid()
	if len(rStateid) == 0 {
		return nil, errors.New("rStateID is required")
	}
	rCountryid := rLocality.GetCountryid()
	if len(rCountryid) == 0 {
		return nil, errors.New("CountryID is required")
	}
	rFax := request.GetFax()
	if len(rFax) == 0 {
		return nil, errors.New("Fax is required")
	}
	rPhonenumber := request.GetPhonenumber()
	if len(rPhonenumber) == 0 {
		return nil, errors.New("PhoneNumber is required")
	}
	rMobile := request.GetMobile()
	if len(rMobile) == 0 {
		return nil, errors.New("Mobile is required")
	}

	rCityID, _ := uuid.FromString(rCityid)
	rStateID, _ := uuid.FromString(rStateid)
	rCountryID, _ := uuid.FromString(rCountryid)

	if rBase != nil {

		rLocationID = uuid.FromStringOrNil(rBase.Id)
		if rLocationID == uuid.Nil {
			rUserID = uuid.FromStringOrNil(rBase.Createdby)
		} else {
			rUserID = uuid.FromStringOrNil(rBase.Updatedby)
			rCreatedUserID = uuid.FromStringOrNil(rBase.Createdby)
		}

		if rBase.Createdatunix != nil {
			createddate = time.Unix(rBase.Createdatunix.Seconds, 0).UTC()
		}

	}

	fmt.Printf("rLocationName = %v \n", rName)
	fmt.Printf("rDisplayName = %v \n", rDisplayname)
	fmt.Printf("rIsactive = %v \n", rIsactive)
	fmt.Printf("rAddressline1 = %v \n", rAddressline1)
	fmt.Printf("rAddressline2 = %v \n", rAddressline2)
	fmt.Printf("rLocality = %+v \n", rLocality)
	fmt.Printf("rCityID = %+v \n", rCityID)
	fmt.Printf("rStateID = %+v \n", rStateID)
	fmt.Printf("rCountryID = %+v \n", rCountryID)
	fmt.Printf("rGloc = %+v \n", rGloc)

	var mGeolocation locationmodel.GeoLocation = locationmodel.GeoLocation{}

	if rGloc != nil {
		mGeolocation = locationmodel.GeoLocation{
			GeoLocationLat:  rGloc.Lat,
			GeoLocationLong: rGloc.Lon,
			Radius:          rGloc.Radius,
		}
	}
	fmt.Printf("location mGeolocation = %+v \n", mGeolocation)

	//Translating into LocationModel Request
	mLocation := &locationmodel.Location{
		Name:         rName,
		DisplayName:  rDisplayname,
		IsActive:     rIsactive,
		AddressLine1: rAddressline1,
		AddressLine2: rAddressline2,
		Fax:          rFax,
		PhoneNumber:  rPhonenumber,
		Mobile:       rMobile,
		GeoLocation:  mGeolocation,
		CityID:       rCityID,
		StateID:      rStateID,
		CountryID:    rCountryID,
	}
	if rLocationID != uuid.Nil {
		//if id not equal to null
		mLocation.ID = rLocationID
		mLocation.CreatedAt = &createddate
		mLocation.UpdatedBy = &rUserID
		mLocation.CreatedBy = &rCreatedUserID
	} else {
		mLocation.CreatedBy = &rUserID
	}

	//lookup db for passed key if no results return empty list
	location, errLocation := r.locationService.CreateLocation(mLocation)

	if errLocation != nil {
		fmt.Printf("error returning from locationService.CreateLocation: %+v \n", errLocation)
	} else {
		errUpdateLocationFeature := r.locationService.UpdateLocationFeatures(location.ID.String(), featureids, rUserID.String())
		if errUpdateLocationFeature != nil {
			fmt.Printf("error returning from locationService.UpdateLocationFeatures: %+v \n", errUpdateLocationFeature)
		}
	}

	var createdatStamp timestamp.Timestamp
	createdatStamp.Seconds = location.CreatedAt.Unix()

	gbase := &gservice.ResponseBase{
		Id:            location.ID.String(),
		Createdatunix: &createdatStamp,
		Createdat:     location.CreatedAt.String(),
		Updatedat:     location.UpdatedAt.String(),
	}

	// returning a New location model that contains base info so that user gets ids only
	glocation := &gservice.Location{B: gbase}

	if location == nil {
		fmt.Printf("locations return from hlocation is nil \n")
	}
	fmt.Printf("db location is =  %+v \n", location)

	return glocation, nil
}

// SaveFeature func
func (r LocationGRPCServer) SaveFeature(ctx context.Context,
	request *gservice.Feature) (*gservice.Feature, error) {
	//getting header to print requestid
	header, _ := metadata.FromIncomingContext(ctx)
	headerRID := header["reqid"]
	fmt.Printf("ReqIDHeader = %+v \n", headerRID)

	//Request Parsing
	var createddate time.Time
	var rFeatureID uuid.UUID
	var rUserID uuid.UUID
	var rCreatedUserID uuid.UUID
	rName := request.GetName()
	if len(rName) == 0 {
		return nil, errors.New("name is required")
	}

	rIsactive := request.GetIsactive()
	rBase := request.GetB()
	rcategorydetailid := uuid.FromStringOrNil(request.GetCategorydetailid())

	if rBase != nil {

		rFeatureID = uuid.FromStringOrNil(rBase.Id)
		if rFeatureID == uuid.Nil {
			rUserID = uuid.FromStringOrNil(rBase.Createdby)
		} else {
			rUserID = uuid.FromStringOrNil(rBase.Updatedby)
			rCreatedUserID = uuid.FromStringOrNil(rBase.Createdby)
		}

		if rBase.Createdatunix != nil {
			createddate = time.Unix(rBase.Createdatunix.Seconds, 0).UTC()
		}

	}

	//Translating into FeatureModel Request
	mFeature := &locationmodel.Feature{
		Name:             rName,
		IsActive:         rIsactive,
		CategoryDetailID: rcategorydetailid,
	}
	if rFeatureID != uuid.Nil {
		//if id not equal to null
		mFeature.ID = rFeatureID
		mFeature.CreatedAt = &createddate
		mFeature.UpdatedBy = &rUserID
		mFeature.CreatedBy = &rCreatedUserID
	} else {
		mFeature.CreatedBy = &rUserID
	}

	//lookup db for passed key if no results return empty list
	feature, errFeature := r.locationService.SaveFeature(mFeature)

	if errFeature != nil {
		fmt.Printf("error returning from locationService.SaveFeature: %+v \n", errFeature)
	}

	var createdatStamp timestamp.Timestamp
	createdatStamp.Seconds = feature.CreatedAt.Unix()

	gbase := &gservice.ResponseBase{
		Id:            feature.ID.String(),
		Createdatunix: &createdatStamp,
		Createdat:     feature.CreatedAt.String(),
		Updatedat:     feature.UpdatedAt.String(),
	}

	// returning a New feature model that contains base info so that user gets ids only
	gfeature := &gservice.Feature{B: gbase}

	if feature == nil {
		fmt.Printf("feature return is nil \n")
	}
	fmt.Printf("db feature is =  %+v \n", feature)

	return gfeature, nil
}

// AllCategoryDetail func
func (r LocationGRPCServer) AllCategoryDetails(ctx context.Context, request *gservice.FilterCategoryDetail) (*gservice.CategoryDetailList, error) {

	//getting header to print requestid
	header, _ := metadata.FromIncomingContext(ctx)
	headerRID := header["reqid"]
	fmt.Printf("ReqIDHeader = %+v \n", headerRID)

	//getting request entity id from responsebase
	rCategoryDetailID := request.GetCategorydetailid()

	fmt.Printf("rCategoryDetailID = %v \n", rCategoryDetailID)

	categoryDetails, errCatDetail := r.locationService.GetCategoryDetail(rCategoryDetailID)
	if errCatDetail != nil {
		fmt.Printf("error returning category detail from locationService.GetCategoryDetail: %+v \n", errCatDetail)
	}
	//fmt.Printf("countries return from hex-location is :%+v \n", countries)
	gcategoryDetail := []*gservice.CategoryDetail{}

	if categoryDetails == nil {
		fmt.Printf("category detail return from hlocation is nil \n")
	}

	for _, c := range *categoryDetails {
		fmt.Printf("city = %+v \n", c)

		var createdatStamp timestamp.Timestamp
		createdatStamp.Seconds = c.CreatedAt.Unix()

		// createdatStamp.set_seconds(c.CreatedAt.Unix())

		gbase := &gservice.ResponseBase{
			Id:            c.ID.String(),
			Createdatunix: &createdatStamp,
			Createdat:     c.CreatedAt.String(),
			Updatedat:     c.UpdatedAt.String()}

		gcategoryDetail = append(gcategoryDetail, &gservice.CategoryDetail{
			Name:             c.Name,
			Categoryid:       c.CategoryID.String(),
			Parentcategoryid: c.ParentCategory.String(),
			Isactive:         c.IsActive,
			B:                gbase})

	}

	reslog := fmt.Sprintf("categorydetail result:  %v", categoryDetails)
	fmt.Println("cities = ", reslog)
	return &gservice.CategoryDetailList{Categorydetails: gcategoryDetail}, nil
}

// SaveCategoryDetail func
func (r LocationGRPCServer) SaveCategoryDetail(ctx context.Context,
	request *gservice.CategoryDetail) (*gservice.CategoryDetail, error) {
	//getting header to print requestid
	header, _ := metadata.FromIncomingContext(ctx)
	headerRID := header["reqid"]
	fmt.Printf("ReqIDHeader = %+v \n", headerRID)

	//Request Parsing
	var createddate time.Time
	var rCatDetailID uuid.UUID
	var rUserID uuid.UUID
	var rCreatedUserID uuid.UUID
	rName := request.GetName()
	if len(rName) == 0 {
		return nil, errors.New("name is required")
	}

	rIsactive := request.GetIsactive()
	rBase := request.GetB()
	rcategoryid := uuid.FromStringOrNil(request.GetCategoryid())

	if rBase != nil {

		rCatDetailID = uuid.FromStringOrNil(rBase.Id)
		if rCatDetailID == uuid.Nil {
			rUserID = uuid.FromStringOrNil(rBase.Createdby)
		} else {
			rUserID = uuid.FromStringOrNil(rBase.Updatedby)
			rCreatedUserID = uuid.FromStringOrNil(rBase.Createdby)
		}

		if rBase.Createdatunix != nil {
			createddate = time.Unix(rBase.Createdatunix.Seconds, 0).UTC()
		}

	}

	//Translating into CategoryDetailModel Request
	mCatDetail := &locationmodel.CategoryDetail{
		Name:       rName,
		IsActive:   rIsactive,
		CategoryID: rcategoryid,
	}
	if rCatDetailID != uuid.Nil {
		//if id not equal to null
		mCatDetail.ID = rCatDetailID
		mCatDetail.CreatedAt = &createddate
		mCatDetail.UpdatedBy = &rUserID
		mCatDetail.CreatedBy = &rCreatedUserID
	} else {
		mCatDetail.CreatedBy = &rUserID
	}

	//lookup db for passed key if no results return empty list
	catDetail, errCatDetail := r.locationService.SaveCategoryDetail(mCatDetail)

	if errCatDetail != nil {
		fmt.Printf("error returning from locationService.SaveCategoryDetail: %+v \n", errCatDetail)
	}

	var createdatStamp timestamp.Timestamp
	createdatStamp.Seconds = catDetail.CreatedAt.Unix()

	gbase := &gservice.ResponseBase{
		Id:            catDetail.ID.String(),
		Createdatunix: &createdatStamp,
		Createdat:     catDetail.CreatedAt.String(),
		Updatedat:     catDetail.UpdatedAt.String(),
	}

	// returning a New categorydetail model that contains base info so that user gets ids only
	gCatDetail := &gservice.CategoryDetail{B: gbase}

	if catDetail == nil {
		fmt.Printf("catDetail return is nil \n")
	}
	fmt.Printf("db catDetail is =  %+v \n", catDetail)

	return gCatDetail, nil
}

// SaveCategory func
func (r LocationGRPCServer) SaveCategory(ctx context.Context,
	request *gservice.Category) (*gservice.Category, error) {
	//getting header to print requestid
	header, _ := metadata.FromIncomingContext(ctx)
	headerRID := header["reqid"]
	fmt.Printf("ReqIDHeader = %+v \n", headerRID)

	//Request Parsing
	var createddate time.Time
	var rCategoryID uuid.UUID
	var rUserID uuid.UUID
	var rCreatedUserID uuid.UUID
	rName := request.GetName()
	if len(rName) == 0 {
		return nil, errors.New("name is required")
	}

	rIsactive := request.GetIsactive()
	rBase := request.GetB()

	if rBase != nil {

		rCategoryID = uuid.FromStringOrNil(rBase.Id)
		if rCategoryID == uuid.Nil {
			rUserID = uuid.FromStringOrNil(rBase.Createdby)
		} else {
			rUserID = uuid.FromStringOrNil(rBase.Updatedby)
			rCreatedUserID = uuid.FromStringOrNil(rBase.Createdby)
		}

		if rBase.Createdatunix != nil {
			createddate = time.Unix(rBase.Createdatunix.Seconds, 0).UTC()
		}

	}

	//Translating into CategoryModel Request
	mCategory := &locationmodel.Category{
		Name:     rName,
		IsActive: rIsactive,
	}
	if rCategoryID != uuid.Nil {
		//if id not equal to null
		mCategory.ID = rCategoryID
		mCategory.CreatedAt = &createddate
		mCategory.UpdatedBy = &rUserID
		mCategory.CreatedBy = &rCreatedUserID
	} else {
		mCategory.CreatedBy = &rUserID
	}

	//lookup db for passed key if no results return empty list
	category, errCategory := r.locationService.SaveCategory(mCategory)

	if errCategory != nil {
		fmt.Printf("error returning from locationService.SaveCategory: %+v \n", errCategory)
	}

	var createdatStamp timestamp.Timestamp
	createdatStamp.Seconds = category.CreatedAt.Unix()

	gbase := &gservice.ResponseBase{
		Id:            category.ID.String(),
		Createdatunix: &createdatStamp,
		Createdat:     category.CreatedAt.String(),
		Updatedat:     category.UpdatedAt.String(),
	}

	// returning a New category model that contains base info so that user gets ids only
	gCategory := &gservice.Category{B: gbase}

	if category == nil {
		fmt.Printf("category return is nil \n")
	}
	fmt.Printf("db category is =  %+v \n", category)

	return gCategory, nil
}
