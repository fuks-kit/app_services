package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	pb "github.com/fuks-kit/app_services/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"slices"
)

var (
	host       = flag.String("host", "localhost:30336", "host:port of gRPC server")
	isInsecure = flag.Bool("insecure", true, "connect without TLS")
)

var validActions = []string{
	"get_events",
	"get_projects",
	"get_karlsruher_transfers",
	"get_links",
}

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()

	actions := flag.Args()
	if len(actions) == 0 {
		actions = validActions
	}

	var opts []grpc.DialOption
	if *host != "" {
		opts = append(opts, grpc.WithAuthority(*host))
	}

	if *isInsecure {
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

	client, err := grpc.Dial(*host, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() { _ = client.Close() }()

	appServices := pb.NewAppServicesClient(client)

	ctx := context.Background()

	//
	// Get Events
	//

	if slices.Contains(actions, "get_events") {
		events, err := appServices.GetEvents(ctx, &emptypb.Empty{})
		if err != nil {
			log.Fatalf("could not get events: %v", err)
		}

		byt, _ := json.MarshalIndent(events, "", "  ")
		log.Printf("GetEvents: %s", byt)
	}

	//
	// Get Projects
	//

	if slices.Contains(actions, "get_projects") {
		projects, err := appServices.GetProjects(ctx, &emptypb.Empty{})
		if err != nil {
			log.Fatalf("could not get events: %v", err)
		}

		byt, _ := json.MarshalIndent(projects, "", "  ")
		log.Printf("GetProjects: %s", byt)
	}

	//
	// Get Kalrsruhe Transfers
	//

	if slices.Contains(actions, "get_karlsruher_transfers") {
		kts, err := appServices.GetKarlsruherTransfers(ctx, &emptypb.Empty{})
		if err != nil {
			log.Fatalf("could not get events: %v", err)
		}

		byt, _ := json.MarshalIndent(kts, "", "  ")
		log.Printf("GetKarlsruherTransfers: %s", byt)
	}

	//
	// Get Links
	//

	if slices.Contains(actions, "get_links") {
		links, err := appServices.GetLinks(ctx, &emptypb.Empty{})
		if err != nil {
			log.Fatalf("could not get events: %v", err)
		}

		byt, _ := json.MarshalIndent(links, "", "  ")
		log.Printf("GetLinks: %s", byt)
	}
}
