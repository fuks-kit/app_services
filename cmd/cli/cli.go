package main

import (
	"context"
	"encoding/json"
	pb "fuks_cloud_services/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	client, err := grpc.Dial("localhost:30336",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() { _ = client.Close() }()

	appServices := pb.NewAppServicesClient(client)

	ctx := context.Background()

	//
	// Get Events
	//

	events, err := appServices.GetEvents(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("could not get events: %v", err)
	}

	byt, _ := json.MarshalIndent(events, "", "  ")
	log.Printf("GetEvents: %s", byt)

	//
	// Get Projects
	//

	projects, err := appServices.GetProjects(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("could not get events: %v", err)
	}

	byt, _ = json.MarshalIndent(projects, "", "  ")
	log.Printf("GetProjects: %s", byt)

	//
	// Get Kalrsruhe Transfers
	//

	kts, err := appServices.GetKarlsruherTransfers(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("could not get events: %v", err)
	}

	byt, _ = json.MarshalIndent(kts, "", "  ")
	log.Printf("GetKarlsruherTransfers: %s", byt)
}
