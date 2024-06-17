package server

import (
	"context"
	"fmt"
	pb "github.com/fuks-kit/app_services/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

var linksCache = newCache[pb.Links]()

const (
	linksTitle = iota
	linksDescription
	linksUrl
)

func (service *AppServices) GetLinks(_ context.Context, _ *emptypb.Empty) (*pb.Links, error) {

	data, validCache := linksCache.get()
	if validCache {
		return data, nil
	}

	readRange := service.config.LinksSheet + "!A2:C"
	resp, err := sheetsService.
		Spreadsheets.
		Values.
		Get(service.config.SheetId, readRange).
		Do()
	if err != nil {
		log.Printf("Unable to retrieve data from sheet: %v", err)
		return nil, fmt.Errorf("unable to retrieve data from sheet")
	}

	var links []*pb.Link
	for _, row := range resp.Values {

		// Ensure the row has at least 3 columns
		if len(row) < 3 {
			row = append(row, make([]interface{}, 3-len(row))...)
		}

		// Get the title and ensure it is a string
		title, ok := row[linksTitle].(string)
		if !ok || title == "" {
			continue
		}

		description, _ := row[linksDescription].(string)

		link, ok := row[linksUrl].(string)
		if !ok || link == "" {
			continue
		}

		links = append(links, &pb.Link{
			Title:       title,
			Description: description,
			Url:         link,
		})
	}

	data = &pb.Links{Items: links}
	linksCache.set(data)

	return data, nil
}
