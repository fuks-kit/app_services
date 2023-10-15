package server

import (
	"context"
	_ "embed"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"log"
)

//go:embed credentials.json
var credentials []byte

var sheetsService *sheets.Service

func init() {
	config, err := google.JWTConfigFromJSON(
		credentials,
		sheets.SpreadsheetsReadonlyScope,
	)
	if err != nil {
		log.Fatalln(err)
	}

	// Don't fuck with this, the mail needs to be set here
	config.Subject = "patrick.zierahn@fuks.org"

	ctx := context.Background()
	ts := config.TokenSource(ctx)

	sheetsService, err = sheets.NewService(ctx, option.WithTokenSource(ts))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}
}
