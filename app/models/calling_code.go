package models

type CallingCode struct {
	CountryCode string   `json:"countryCode"`
	Flag        string   `json:"flag"`
	Root        string   `json:"root"`
	Suffixes    []string `json:"suffixes"`
}
