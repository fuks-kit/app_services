package server

import (
	"context"
	_ "embed"
	auth "golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"log"
	"os"
)

const credentialsFile = "credentials.json"

var sheetsService *sheets.Service

func init() {
	ctx := context.Background()

	params := auth.CredentialsParams{
		Scopes: []string{
			sheets.SpreadsheetsReadonlyScope,
		},
		//Subject: "patrick.zierahn@fuks.org",
		Subject: "fcs-account@fuks-app.iam.gserviceaccount.com",
	}

	var credentials *auth.Credentials

	if _, exist := os.Stat(credentialsFile); exist == nil {
		//
		// Use local credentials
		//

		log.Printf("Using local credentials: %s", credentialsFile)

		jsonKey, err := os.ReadFile(credentialsFile)
		if err != nil {
			log.Fatalln(err)
		}

		cred, err := auth.CredentialsFromJSONWithParams(ctx, jsonKey, params)
		if err != nil {
			log.Fatalln(err)
		}

		credentials = cred

	} else {
		cred, err := auth.FindDefaultCredentialsWithParams(ctx, params)
		if err != nil {
			log.Fatalln(err)
		}

		credentials = cred
	}

	sheet, err := sheets.NewService(ctx,
		option.WithCredentials(credentials),
	)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	sheetsService = sheet
}
