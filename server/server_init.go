package server

import (
	"context"
	_ "embed"
	"encoding/json"
	auth "golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"log"
	"os"
)

var sheetsService *sheets.Service

type Debug struct {
	Type                    string `json:"type"`
	ProjectId               string `json:"project_id"`
	ClientEmail             string `json:"client_email"`
	ClientId                string `json:"client_id"`
	AuthUri                 string `json:"auth_uri"`
	TokenUri                string `json:"token_uri"`
	AuthProviderX509CertUrl string `json:"auth_provider_x509_cert_url"`
	ClientX509CertUrl       string `json:"client_x509_cert_url"`
	UniverseDomain          string `json:"universe_domain"`
}

func init() {
	ctx := context.Background()

	params := auth.CredentialsParams{
		Scopes: []string{
			sheets.SpreadsheetsReadonlyScope,
		},
		Subject: "patrick.zierahn@fuks.org",
	}

	log.Printf("######## GOOGLE_APPLICATION_CREDENTIALS: %s", os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))

	credentials, err := auth.FindDefaultCredentialsWithParams(ctx, params)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("######## ProjectID: %s", credentials.ProjectID)

	var deb Debug
	err = json.Unmarshal(credentials.JSON, &deb)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("########: deb=%+v", deb)

	sheet, err := sheets.NewService(ctx,
		option.WithCredentials(credentials),
	)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	sheetsService = sheet
}
