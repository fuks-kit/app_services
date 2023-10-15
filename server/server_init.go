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

	log.SetFlags(log.LstdFlags | log.Lshortfile)

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

	token, _ := credentials.TokenSource.Token()

	log.Printf("######## credentials.ProjectID: %s", credentials.ProjectID)
	log.Printf("######## JSON len: %d", len(credentials.JSON))
	log.Printf("######## tok: %d %v", len(token.AccessToken), token.Expiry)
	log.Printf("######## token.TokenType: %v", token.TokenType)

	config, err := auth.JWTConfigFromJSON(
		[]byte(token.AccessToken),
		sheets.SpreadsheetsReadonlyScope,
	)
	if err != nil {
		log.Fatalln(err)
	}

	// Don't fuck with this. The email to the service account is required here!
	config.Subject = "patrick.zierahn@fuks.org"

	ts := config.TokenSource(ctx)

	sheet, err := sheets.NewService(ctx,
		//option.WithCredentials(credentials),
		option.WithTokenSource(ts),
	)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	sheetsService = sheet
}
