package meldeschein_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetMeldeschein(t *testing.T) {
	req := client.NewGetMeldeschein()
	// req.RequestBody().Identifikation.Erzeugung = "2020-02-02"
	req.RequestBody().Identifikation.Schnittstelle = "GUESTLINE_OMNIBOOST"
	req.RequestBody().AnfrageDaten.Meldescheinnummer = "39899"
	// req.RequestBody().AnfrageDaten.Buchungsnummer = "1223"
	// req.RequestBody().AnfrageDaten.OrtID = 363
	// req.RequestBody().ArrivalDate = guestline.DateTime{time.Now().AddDate(0, 0, -7)}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		t.Error(err)
	}

	log.Println(string(b))
}
