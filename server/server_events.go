package server

import (
	"context"
	"fmt"
	pb "fuks_cloud_services/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
)

func (service *AppServices) GetEvents(_ context.Context, _ *emptypb.Empty) (*pb.Events, error) {
	readRange := service.config.EventsSheet + "!A2:H"

	resp, err := sheetsService.
		Spreadsheets.
		Values.
		Get(service.config.SheetId, readRange).
		Do()
	if err != nil {
		log.Printf("Unable to retrieve data from sheet: %v", err)
		return nil, fmt.Errorf("unable to retrieve data from sheet")
	}

	var events []*pb.Event
	for _, row := range resp.Values {

		title, ok := row[0].(string)
		if !ok || title == "" {
			continue
		}

		subtitle, ok := row[1].(string)
		if !ok || subtitle == "" {
			continue
		}

		dateStr, ok := row[2].(string)
		if !ok || dateStr == "" {
			continue
		}

		timeStr, ok := row[3].(string)
		if !ok || timeStr == "" {
			continue
		}

		loc, _ := time.LoadLocation("Europe/Berlin")
		eventDate := dateStr + " " + timeStr
		date, err := time.ParseInLocation("02/01/2006 15:04", eventDate, loc)
		if err != nil {
			log.Printf("Unable to parse date: %v", err)
			continue
		}

		log.Printf("Date: %v", date)

		location, _ := row[4].(string)
		contactName, _ := row[5].(string)
		contactEMail, _ := row[6].(string)
		contactImage, _ := row[7].(string)

		event := &pb.Event{
			Title:    title,
			Subtitle: subtitle,
			Date:     timestamppb.New(date),
			Location: location,
			Contact: &pb.Contact{
				Name:     contactName,
				EMail:    contactEMail,
				ImageUrl: contactImage,
			},
		}

		events = append(events, event)
	}

	return &pb.Events{Items: events}, nil
}
