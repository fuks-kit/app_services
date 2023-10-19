package server

import (
	"context"
	"fmt"
	pb "github.com/fuks-kit/app_services/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (service *AppServices) GetKarlsruherTransfers(_ context.Context, _ *emptypb.Empty) (*pb.KarlsruherTransfers, error) {
	readRange := service.config.KTSheet + "!A2:H"

	resp, err := sheetsService.
		Spreadsheets.
		Values.
		Get(service.config.SheetId, readRange).
		Do()
	if err != nil {
		log.Printf("Unable to retrieve data from sheet: %v", err)
		return nil, fmt.Errorf("unable to retrieve data from sheet")
	}

	var kts []*pb.KarlsruherTransfer
	for _, row := range resp.Values {

		title, ok := row[0].(string)
		if !ok || title == "" {
			continue
		}

		subtitle, ok := row[1].(string)
		if !ok || subtitle == "" {
			continue
		}

		preview, ok := row[2].(string)
		if !ok || preview == "" {
			continue
		}

		pdf, ok := row[3].(string)
		if !ok || pdf == "" {
			continue
		}

		kt := &pb.KarlsruherTransfer{
			Title:    title,
			Subtitle: subtitle,
			ImageUrl: preview,
			PdfUrl:   pdf,
		}

		kts = append(kts, kt)
	}

	return &pb.KarlsruherTransfers{Items: kts}, nil
}
