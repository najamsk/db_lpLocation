package main

import (
	"fmt"
	"log"

	//"github.com/gogo/protobuf/proto"
	"flag"
	"os"

	gservice "github.com/launchpad7/lp7_micro_location/grpc"
	"golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// TODO: this should be rename to locationGRPCServer

func main() {

	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "missing subcommand: list or add")
		os.Exit(1)
	}

	conn, err := grpc.Dial(":8888", grpc.WithInsecure())

	if err != nil {
		fmt.Fprintf(os.Stderr, "could not connect to backend: %v\n", err)
		os.Exit(1)
	}
	client := gservice.NewLocationRPCClient(conn)

	gclient := NewLocationGRPCClient()

	if gclient == nil {
		log.Fatal("our gclient that is calling grpc server is not valid object")
	}

	switch cmd := flag.Arg(0); cmd {
	case "countries_filter":
		err = gclient.ListCountriesByFilter(context.Background(), client)
	case "states_filter":
		err = gclient.ListStatesByFilter(context.Background(), client)
	case "cities_filter":
		err = gclient.ListCitiesByFilter(context.Background(), client)
	case "locations_filter":
		err = gclient.ListLocationsByFilter(context.Background(), client)
	case "categories_filter":
		err = gclient.ListCategoriesByFilter(context.Background(), client)
	case "catsubs_filter":
		err = gclient.ListCategoryDetailsByFilter(context.Background(), client)
	case "features_filter":
		err = gclient.ListFeaturesByFilter(context.Background(), client)
	case "staff_filter":
		err = gclient.ListStaffByFilter(context.Background(), client)
	case "location_create":
		err = gclient.CreateLocation(context.Background(), client)
	default:
		err = fmt.Errorf("unknown subcommand %s", cmd)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
