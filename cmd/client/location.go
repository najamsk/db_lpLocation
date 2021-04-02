package main

import (
	"fmt"
	//"github.com/gogo/protobuf/proto"
	gservice "github.com/launchpad7/lp7_micro_location/grpc"
	"golang.org/x/net/context"

	//grpc "google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

//LocationGRPCClient will have members funcs on it
type LocationGRPCClient struct {
}

//NewLocationGRPCClient will give you struct
func NewLocationGRPCClient() *LocationGRPCClient {
	return &LocationGRPCClient{}
}

//ListCountriesByFilter will do
func (LocationGRPCClient) ListCountriesByFilter(ctx context.Context, client gservice.LocationRPCClient) error {
	// new request context
	ctx22 := context.Background()
	ctx22 = metadata.NewOutgoingContext(ctx22,
		metadata.Pairs("dev", "habibi", "network", "moftak ka net", "ReqID", "345"),
	)

	l, err := client.AllCountries(ctx22, &gservice.FilterLocation{})

	if err != nil {
		return fmt.Errorf("could not add task in the backend: %v", err)
	}

	res := l.GetCountries()
	fmt.Println("countries result is = ", res)
	return nil

}

//ListStatesByFilter will do
func (LocationGRPCClient) ListStatesByFilter(ctx context.Context, client gservice.LocationRPCClient) error {
	// new request context
	ctx22 := context.Background()
	ctx22 = metadata.NewOutgoingContext(ctx22, metadata.Pairs("dev", "habibi", "network", "moftak ka net", "ReqID", "345"))

	l, err := client.AllStates(ctx22, &gservice.FilterLocation{})
	//l, err := client.StatesByCountry(ctx22, &gservice.Country{B: &gservice.ResponseBase{}})

	if err != nil {
		//return fmt.Errorf("%v", err)
		return err
	}

	res := l.GetStates()
	fmt.Println("statelist result is = ", res)
	return nil

}

//ListCitiesByFilter will do
func (LocationGRPCClient) ListCitiesByFilter(ctx context.Context, client gservice.LocationRPCClient) error {
	// new request context
	ctx22 := context.Background()
	ctx22 = metadata.NewOutgoingContext(ctx22, metadata.Pairs("dev", "habibi", "network", "moftak ka net", "ReqID", "345"))

	l, err := client.AllCities(ctx22, &gservice.FilterLocation{})
	//l, err := client.StatesByCountry(ctx22, &gservice.Country{B: &gservice.ResponseBase{}})

	if err != nil {
		//return fmt.Errorf("%v", err)
		return err
	}

	res := l.GetCities()
	fmt.Println("citylist result is = ", res)
	return nil

}

//ListLocationsByFilter will do
func (LocationGRPCClient) ListLocationsByFilter(ctx context.Context, client gservice.LocationRPCClient) error {
	// new request context
	ctx22 := context.Background()
	ctx22 = metadata.NewOutgoingContext(ctx22, metadata.Pairs("dev", "habibi", "network", "moftak ka net", "ReqID", "345"))

	l, err := client.AllLocations(ctx22, &gservice.FilterLocation{Cityid: "d1b0ebbe-55e0-409c-a80c-13d8e5125cfa",
		Stateid:    "1257bf8c-3971-4455-903a-aa5fd2990450",
		Countryid:  "",
		Locationid: "060ebd73-aa14-48ba-b5c3-71653715319e"})
	//l, err := client.StatesByCountry(ctx22, &gservice.Country{B: &gservice.ResponseBase{}})

	if err != nil {
		//return fmt.Errorf("%v", err)
		return err
	}

	res := l.GetLocations()
	fmt.Println("LocationsList result is = ", res)
	return nil

}

//ListCategoriesByFilter will do
func (LocationGRPCClient) ListCategoriesByFilter(ctx context.Context, client gservice.LocationRPCClient) error {
	// new request context
	ctx22 := context.Background()
	ctx22 = metadata.NewOutgoingContext(ctx22, metadata.Pairs("dev", "habibi", "network", "moftak ka net", "ReqID", "345"))

	l, err := client.AllCategories(ctx22,
		&gservice.FilterCategory{
			Categoryid: "d1b0ebbe-55e0-409c-a80c-13d8e5125cfa",
		})
	//l, err := client.StatesByCountry(ctx22, &gservice.Country{B: &gservice.ResponseBase{}})

	if err != nil {
		//return fmt.Errorf("%v", err)
		return err
	}

	res := l
	fmt.Println("CategoryList result is = ", res)
	return nil

}

//ListCategoryDetailsByFilter will do
func (LocationGRPCClient) ListCategoryDetailsByFilter(ctx context.Context, client gservice.LocationRPCClient) error {
	// new request context
	ctx22 := context.Background()
	ctx22 = metadata.NewOutgoingContext(ctx22, metadata.Pairs("dev", "habibi", "network", "moftak ka net", "ReqID", "345"))

	l, err := client.AllCategoryDetails(ctx22,
		&gservice.FilterCategoryDetail{
			Categoryid: "d1b0ebbe-55e0-409c-a80c-13d8e5125cfa",
		})
	//l, err := client.StatesByCountry(ctx22, &gservice.Country{B: &gservice.ResponseBase{}})

	if err != nil {
		//return fmt.Errorf("%v", err)
		return err
	}

	res := l
	fmt.Println("CategoryList result is = ", res)
	return nil

}

//ListFeaturesByFilter will do
func (LocationGRPCClient) ListFeaturesByFilter(ctx context.Context, client gservice.LocationRPCClient) error {
	// new request context
	ctx22 := context.Background()
	ctx22 = metadata.NewOutgoingContext(ctx22, metadata.Pairs("dev", "habibi", "network", "moftak ka net", "ReqID", "345"))

	l, err := client.AllFeatures(ctx22,
		&gservice.FilterFeature{
			Categorydetailid: "d1b0ebbe-55e0-409c-a80c-13d8e5125cfa",
		})
	//l, err := client.StatesByCountry(ctx22, &gservice.Country{B: &gservice.ResponseBase{}})

	if err != nil {
		//return fmt.Errorf("%v", err)
		return err
	}

	res := l
	fmt.Println("CategoryList result is = ", res)
	return nil

}

//ListStaffByFilter will do
func (LocationGRPCClient) ListStaffByFilter(ctx context.Context, client gservice.LocationRPCClient) error {
	// new request context
	ctx22 := context.Background()
	ctx22 = metadata.NewOutgoingContext(ctx22, metadata.Pairs("dev", "habibi", "network", "moftak ka net", "ReqID", "345"))

	l, err := client.AllStaff(ctx22,
		&gservice.FilterStaff{
			Userid: "d1b0ebbe-55e0-409c-a80c-13d8e5125cfa",
		})
	//l, err := client.StatesByCountry(ctx22, &gservice.Country{B: &gservice.ResponseBase{}})

	if err != nil {
		//return fmt.Errorf("%v", err)
		return err
	}

	res := l
	fmt.Println("Stafflist result is = ", res)
	return nil

}

//CreateLocation will do
func (LocationGRPCClient) CreateLocation(ctx context.Context, client gservice.LocationRPCClient) error {
	// new request context
	ctx22 := context.Background()
	ctx22 = metadata.NewOutgoingContext(ctx22, metadata.Pairs("dev", "habibi", "network", "moftak ka net", "ReqID", "345"))

	//setup
	glocality := &gservice.Locality{
		Countryid: "b4a029f4-f4f2-4578-9ce4-52dc226b2d95",
		Stateid:   "1b8d1369-7026-45d1-b47e-7a5a4d37f02f",
		Cityid:    "552b164d-f6b3-4249-ad0f-0789e0624e98",
	}

	gGeolocation := &gservice.GeoLocation{
		Lat:    55.55,
		Lon:    66.66,
		Radius: 3.3,
	}

	fmt.Printf("location geoloc= %+v \n", gGeolocation)

	l, err := client.SaveLocation(ctx22, &gservice.Location{
		Name:        "Naqli moftak",
		Displayname: "naqli mofatk he asli hai bahi.",
		Isactive:    true,
		Locality:    glocality,
		Phonenumber: "046-90000",
		Mobile:      "034500012",
		Fax:         "0990-123123",
		//Gloc:        gGeolocation,
	})

	if err != nil {
		//return fmt.Errorf("%v", err)
		return err
	}

	res := l
	fmt.Println("LocationsList result is = ", res)
	return nil

}
