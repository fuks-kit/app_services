package server

import (
	"context"
	_ "embed"
	auth "golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"log"
)

var sheetsService *sheets.Service

func init() {
	ctx := context.Background()

	params := auth.CredentialsParams{
		Scopes: []string{
			sheets.SpreadsheetsReadonlyScope,
		},
		Subject: "113222746783594416669",
	}

	credentials, err := auth.FindDefaultCredentialsWithParams(ctx, params)
	if err != nil {
		log.Fatalln(err)
	}

	sheet, err := sheets.NewService(ctx,
		option.WithCredentials(credentials),
	)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	sheetsService = sheet
}
