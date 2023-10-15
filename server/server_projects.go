package server

import (
	"context"
	"fmt"
	pb "fuks_cloud_services/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (service *AppServices) GetProjects(_ context.Context, _ *emptypb.Empty) (*pb.Projects, error) {
	readRange := service.config.ProjectsSheet + "!A2:G"

	resp, err := sheetsService.
		Spreadsheets.
		Values.
		Get(service.config.SheetId, readRange).
		Do()
	if err != nil {
		log.Printf("Unable to retrieve data from sheet: %v", err)
		return nil, fmt.Errorf("unable to retrieve data from sheet")
	}

	var projects []*pb.Project
	for _, row := range resp.Values {

		title, ok := row[0].(string)
		if !ok || title == "" {
			continue
		}

		subtitle, ok := row[1].(string)
		if !ok || subtitle == "" {
			continue
		}

		label, _ := row[2].(string)

		managerName, ok := row[3].(string)
		if !ok || managerName == "" {
			continue
		}

		managerEmail, ok := row[4].(string)
		if !ok || managerEmail == "" {
			continue
		}

		managerImage, _ := row[5].(string)

		details, ok := row[6].(string)
		if !ok || details == "" {
			continue
		}

		projects = append(projects, &pb.Project{
			Title:    title,
			Subtitle: subtitle,
			Label:    label,
			Details:  details,
			Manager: &pb.Contact{
				Name:     managerName,
				EMail:    managerEmail,
				ImageUrl: managerImage,
			},
		})
	}

	return &pb.Projects{Items: projects}, nil
}