package workspace

import (
	"log"
)

const (
	sheetId     = "1lg1VTtVhQ8gFoe08L0DPGUaJthlb4lvID0GhAkDt7sY"
	eventsSheet = "Events"
)

func ReadSheet() {
	readRange := eventsSheet + "!A2:G"

	resp, err := sheetsService.
		Spreadsheets.
		Values.
		Get(sheetId, readRange).
		Do()
	if err != nil {
		log.Fatal(err)
	}

	if len(resp.Values) == 0 {
		log.Println("No data found.")
		return
	}

	for _, row := range resp.Values {
		log.Printf("%s, %s, %s, %s, %s, %s\n", row[0], row[1], row[2], row[3], row[4], row[5])
	}
}
