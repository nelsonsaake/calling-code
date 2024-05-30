package models

type CallingCode struct {
	CountryCode string `json:"countryCode"`
	CallingCode string `json:"callingCode"`
	Flag        string `json:"flag"`
}
