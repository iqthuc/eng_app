package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func GetGoogleSheetData(w http.ResponseWriter, r *http.Request) {
	const (
		spreadsheetID = "1Vz-LjgI-ha13k368d5hvvyD-thmH2dJ2CdCfn65rwPo"
		readRange     = "Main!A:F"
		credentials   = "iqthuc-cccd7defcaf3.json"
	)
	var sheetsService *sheets.Service
	creds, err := os.ReadFile(credentials)
	if err != nil {
		log.Fatalf("Unable to read credentials file: %v", err)
	}

	config, err := google.JWTConfigFromJSON(creds, sheets.SpreadsheetsScope)
	if err != nil {
		log.Fatalf("Unable to create JWT config: %v", err)
	}

	client := config.Client(context.Background())
	sheetsService, err = sheets.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to create Google Sheets service: %v", err)
	}

	resp, err := sheetsService.Spreadsheets.Values.Get(spreadsheetID, readRange).Context(r.Context()).Do()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	vocabs := ConvertToVocabularySlice(resp.Values)
	data, _ := json.Marshal(vocabs)
	
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

type Vocabulary struct {
	Word     string `json:"word"`
	Phonetic string `json:"phonetic"`
	Meaning  string `json:"meaning"`
	Example  string `json:"example"`
}

func ConvertToVocabularySlice(interfaceSlice [][]interface{}) []Vocabulary {
	var vocabularySlice []Vocabulary

	for _, entry := range interfaceSlice {
		var word, phonetic, meaning, example string

		if len(entry) > 0 {
			word = fmt.Sprintf("%v", entry[0])
		}
		if len(entry) > 1 {
			phonetic = fmt.Sprintf("%v", entry[1])
		}
		if len(entry) > 2 {
			meaning = fmt.Sprintf("%v", entry[2])
		}
		if len(entry) > 3 {
			example = fmt.Sprintf("%v", entry[3])
		}

		vocabulary := Vocabulary{
			Word:     word,
			Phonetic: phonetic,
			Meaning:  meaning,
			Example:  example,
		}
		vocabularySlice = append(vocabularySlice, vocabulary)
	}

	return vocabularySlice
}
