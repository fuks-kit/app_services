package server

import (
	"context"
	"fmt"
	pb "github.com/fuks-kit/app_services/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"sort"
	"time"
)

var eventsCache = newCache[pb.Events]()

func (service *AppServices) GetEvents(_ context.Context, _ *emptypb.Empty) (*pb.Events, error) {

	data, validCache := eventsCache.get()
	if validCache {
		return data, nil
	}

	readRange := service.config.EventsSheet + "!A2:K"
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

		if len(row) < 11 {
			row = append(row, make([]interface{}, 11-len(row))...)
		}

		title, ok := row[0].(string)
		if !ok || title == "" {
			continue
		}

		details, ok := row[1].(string)
		if !ok || details == "" {
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

		loc, err := time.LoadLocation("Europe/Berlin")
		if err != nil {
			log.Printf("Unable to load location: %v", err)
			continue
		}

		eventDate := dateStr + " " + timeStr
		date, err := time.ParseInLocation("2/1/2006 15:04", eventDate, loc)
		if err != nil {
			log.Printf("Unable to parse date: %v", err)
			continue
		}

		location, _ := row[4].(string)
		contactName, _ := row[5].(string)
		contactEMail, _ := row[6].(string)
		contactImage, _ := row[7].(string)
		label, _ := row[8].(string)
		buttonText, _ := row[9].(string)
		buttonHref, _ := row[10].(string)

		event := &pb.Event{
			Title:    title,
			Details:  details,
			Date:     timestamppb.New(date),
			Location: location,
			Contact: &pb.Contact{
				Name:     contactName,
				EMail:    contactEMail,
				ImageUrl: contactImage,
			},
			Label:      label,
			ButtonText: buttonText,
			ButtonHref: buttonHref,
		}

		events = append(events, event)
	}

	// Sort events by date
	sort.Slice(events, func(i, j int) bool {
		return events[i].Date.AsTime().Before(events[j].Date.AsTime())
	})

	data = &pb.Events{Items: events}
	eventsCache.set(data)

	return data, nil
}
