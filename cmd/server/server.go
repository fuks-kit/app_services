package main

import (
	pb "github.com/fuks-kit/app_services/proto"
	"github.com/fuks-kit/app_services/server"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	service := server.NewAppServices(server.Config{
		SheetId:       os.Getenv("FUKS_APP_CONTENT_SHEET_ID"),
		EventsSheet:   "Events",
		ProjectsSheet: "Projects",
		KTSheet:       "Karlsruher Transfer",
		LinksSheet:    "Links",
	})

	grpcServer := grpc.NewServer()
	pb.RegisterAppServicesServer(grpcServer, service)

	port := os.Getenv("PORT")
	if port == "" {
		port = "30336"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
