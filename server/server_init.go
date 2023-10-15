package server

import (
	"context"
	_ "embed"
	"google.golang.org/api/sheets/v4"
	"log"
)

var sheetsService *sheets.Service

func init() {
	ctx := context.Background()

	//params := auth.CredentialsParams{
	//	Scopes: []string{
	//		sheets.SpreadsheetsReadonlyScope,
	//	},
	//	Subject: "patrick.zierahn@fuks.org",
	//}
	//
	//credentials, err := auth.FindDefaultCredentialsWithParams(ctx, params)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	sheet, err := sheets.NewService(ctx) //option.WithCredentials(credentials),

	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	sheetsService = sheet
}
