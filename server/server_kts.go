package server

import (
	"context"
	"fmt"
	pb "github.com/fuks-kit/app_services/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"sync"
	"time"
)

var (
	ktMutex sync.RWMutex
	ktCache *pb.KarlsruherTransfers
	ktTime  time.Time
)

func (service *AppServices) GetKarlsruherTransfers(_ context.Context, _ *emptypb.Empty) (*pb.KarlsruherTransfers, error) {

	ktMutex.RLock()
	validCache := ktCache != nil && time.Now().Sub(ktTime) < 5*time.Minute
	ktMutex.RUnlock()

	if validCache {
		return ktCache, nil
	}

	ktMutex.Lock()
	defer ktMutex.Unlock()

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

		if len(row) < 4 {
			row = append(row, make([]interface{}, 4-len(row))...)
		}

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

	ktCache = &pb.KarlsruherTransfers{Items: kts}
	ktTime = time.Now()

	return ktCache, nil
}
