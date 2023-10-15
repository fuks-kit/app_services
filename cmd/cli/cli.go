package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	pb "fuks_cloud_services/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

const (
	host       = "app-services-befklaxdqa-ey.a.run.app:443"
	isInsecure = false
	//host       = "localhost:30336"
	//isInsecure = true
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var opts []grpc.DialOption
	if host != "" {
		opts = append(opts, grpc.WithAuthority(host))
	}

	if isInsecure {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		// Note: On the Windows platform, use of x509.SystemCertPool() requires
		// go version 1.18 or higher.
		systemRoots, err := x509.SystemCertPool()
		if err != nil {
			log.Fatalf("failed to read system root certificate pool: %s", err)
		}

		cred := credentials.NewTLS(&tls.Config{
			RootCAs: systemRoots,
		})

		opts = append(opts, grpc.WithTransportCredentials(cred))
	}

	client, err := grpc.Dial(host, opts...)
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
