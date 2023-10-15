package server

import (
	pb "fuks_cloud_services/proto"
)

type Config struct {
	SheetId       string
	EventsSheet   string
	ProjectsSheet string
	KTSheet       string
}

type AppServices struct {
	pb.AppServicesServer
	config Config
}

func NewAppServices(config Config) *AppServices {
	return &AppServices{config: config}
}
