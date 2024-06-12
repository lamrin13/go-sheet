package internal

import (
	"context"
	"log"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type sheet struct {
	Srv        *sheets.Service
	SheetId    string
	SheetRange string
}

func getService() (*sheet, error) {
	ctx := context.Background()
	keyFile := "key.json"

	// Read the service account key file
	data, err := os.ReadFile(keyFile)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// Authenticate using the service account key file
	config, err := google.JWTConfigFromJSON(data, sheets.SpreadsheetsScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	client := config.Client(ctx)

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}
	return &sheet{
		Srv: srv,
	}, nil

}
