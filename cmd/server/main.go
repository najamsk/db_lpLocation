package main

import (
	"fmt"
	//"github.com/gogo/protobuf/proto"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	gservice "github.com/launchpad7/lp7_micro_location/grpc"
	"github.com/launchpad7/lp7_micro_location/location"
	"github.com/launchpad7/lp7_micro_location/repository/cockroachdb"
	grpc "google.golang.org/grpc"
)

func main() {
	// working
	grpcServer := grpc.NewServer()
	cockroachRepo, errRepo := cockroachdb.NewCockroachRepository()

	if errRepo != nil {
		fmt.Println("cockroachRepo throws error")
	}

	defer cockroachRepo.ClientClose()

	locationService := location.NewLocationService(cockroachRepo)
	locgrpcserver := NewLocationGRPCServer(locationService)
	gservice.RegisterLocationRPCServer(grpcServer, locgrpcserver)
	l, err := net.Listen("tcp", ":8888")

	if err != nil {
		log.Fatalf("could not listen to port :8888 %v", err)
	}
	//fmt.Println("listening on port :8888")
	//log.Fatal(grpcServer.Serve(l))
	errs := make(chan error, 2)
	go func() {
		fmt.Println("Listening on port :8888")
		//errs <- http.ListenAndServe(httpPort(), r)
		errs <- grpcServer.Serve(l)

	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()
	fmt.Printf("Terminated %s", <-errs)

}
