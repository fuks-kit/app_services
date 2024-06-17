package server

import (
	pb "github.com/fuks-kit/app_services/proto"
)

type Config struct {
	SheetId       string
	EventsSheet   string
	ProjectsSheet string
	KTSheet       string
	LinksSheet    string
}

type AppServices struct {
	pb.AppServicesServer
	config Config
}

func NewAppServices(config Config) *AppServices {
	return &AppServices{config: config}
}
