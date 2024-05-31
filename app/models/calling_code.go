package models

type CallingCode struct {
	CountryName string   `json:"countryName"`
	CountryCode string   `json:"countryCode"`
	Flag        string   `json:"flag"`
	Root        string   `json:"root"`
	Suffixes    []string `json:"suffixes"`
}
