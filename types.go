package meldeschein

import (
	"fmt"

	"github.com/hashicorp/go-multierror"
)

type Identifikation struct {
	Erzeugung     Date   `xml:"erzeugung"`
	Schnittstelle string `xml:"schnittstelle"`
	Kurverwaltung int    `xml:"kurverwaltung"`
	BenutzerID    int    `xml:"benutzerid"`
	Verarbeitung  string `xml:"verarbeitung"`
}

type Fehlermeldungen []Fehler

func (ff Fehlermeldungen) Error() string {
	errs := []error{}
	for _, f := range ff {
		errs = append(errs, error(f))
	}
	return multierror.ListFormatFunc(errs)
}

type Fehler struct {
	Code         string `xml:"code"`
	Beschreibung string `xml:"beschreibung"`
	Bezug        string `xml:"bezug"`
}

func (f Fehler) Error() string {
	return fmt.Sprintf("%s: %s (%s)", f.Bezug, f.Beschreibung, f.Code)
}

type Meldeschein struct {
	Buchungsnummer    string `xml:"buchungsnummer,omitempty"`
	Meldescheinnummer int    `xml:"meldescheinnummer,omitempty"`
	FirmaID           int    `xml:"firmaid,omitempty"`
	ObjektID          int    `xml:"objektid,omitempty"`
	Anreise           Date   `xml:"anreise"`
	Abreise           Date   `xml:"abreise"`
	KategorieID       int    `xml:"kategorieid,omitempty"`
	AnredeID          int    `xml:"anredeid,omitempty"`
	Name              string `xml:"name"`
	Vorname           string `xml:"vorname"`
	Strasse           string `xml:"strasse"`
	Hausnummer        string `xml:"hausnummer"`
	Plz               string `xml:"plz"`
	Ort               string `xml:"ort"`
	LandID            int    `xml:"landid"`
	// Land                   Land   `xml:"land"`
	StaatsangehoerigkeitID int `xml:"staatsangehoerigkeitid"`
	// Staatsangehoerigkeit   Staatsangehoerigkeit `xml:"staatsangehoerigkeit"`
	WeitereAngaben string `xml:"weitere_angaben"`
	Ausweisnr      string `xml:"ausweisnr"`
	Kfzkennzeichen string `xml:"kfzkennzeichen"`
	Geburtsdatum   Date   `xml:"geburtsdatum"`
	Begleitperson  []struct {
		AnredeID     int    `xml:"anredeid"`
		Name         string `xml:"name"`
		Vorname      string `xml:"vorname"`
		Geburtsdatum string `xml:"geburtsdatum"`
		KategorieID  int    `xml:"kategorieid"`
		Anreise      Date   `xml:"anreise"`
		Abreise      Date   `xml:"abreise"`
		Ausweisnr    string `xml:"ausweisnr"`
	} `xml:"begleitperson"`
}

type Land struct {
	Text    string `xml:",chardata"`
	LandID  string `xml:"land-id"`
	LandIso string `xml:"land-iso"`
}

type Staatsangehoerigkeit struct {
	Text    string `xml:",chardata"`
	LandID  string `xml:"land-id"`
	LandIso string `xml:"land-iso"`
}

type HoleMeldeschein struct {
	OrtID             string `xml:"ort-id"`
	Buchungsnummer    string `xml:"buchungsnummer"`
	Meldescheinnummer int    `xml:"meldescheinnummer"`
	Firmaid           string `xml:"firmaid"`
	Objektid          string `xml:"objektid"`
	Hauptperson       struct {
		Personid string `xml:"personid"`
		Adresse  struct {
			Adressartid string `xml:"adressartid"`
			Vorname     string `xml:"vorname"`
			Name        string `xml:"name"`
			Land        struct {
				LandID  string `xml:"land-id"`
				LandIso string `xml:"land-iso"`
			} `xml:"land"`
		} `xml:"adresse"`
		Geburtsdatum           string `xml:"geburtsdatum"`
		Staatsangehoerigkeitid struct {
			LandID  string `xml:"land-id"`
			LandIso string `xml:"land-iso"`
		} `xml:"staatsangehoerigkeitid"`
		Kategorieid string `xml:"kategorieid"`
		Anreise     string `xml:"anreise"`
		Abreise     string `xml:"abreise"`
		Barcode     string `xml:"barcode"`
	} `xml:"hauptperson"`
	Betrag             string `xml:"betrag"`
	Abrechnungstatusid string `xml:"abrechnungstatusid"`
}
