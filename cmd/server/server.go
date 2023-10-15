package main

import (
	"fuks_cloud_services/workspace"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	workspace.ReadSheet()
}
