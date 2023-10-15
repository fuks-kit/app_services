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

	events, err := appServices.GetEvents(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("could not get events: %v", err)
	}

	byt, _ := json.MarshalIndent(events, "", "  ")
	log.Printf("GetEvents: %s", byt)
}