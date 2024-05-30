package main

import (
	"net/http"

	"github.com/nelsonsaake/go-ns/axios"
	"github.com/nelsonsaake/go-ns/pretty"
)

type CallingCode struct {
	CountryCode string `json:"countryCode"`
	CallingCode string `json:"callingCode"`
	Flag        string `json:"flag"`
}

type Country struct {
	CountryCode string `json:"cca2"`
	Flags       struct {
		Png string `json:"flags.png"`
	} `json:"flags"`
	Idd struct {
		Root     string   `json:"root"`
		Suffixes []string `json:"suffixes"`
	} `json:"idd"`
}

func (country *Country) callingCode() CallingCode {
	return CallingCode{
		Flag: country.Flags.Png,
		// CountryCode: country.Name.,
	}
}

func main1() {

	var (
		url = "https://restcountries.com/v3.1/all?fields=name,flags,idd"
	)

	client := axios.New()

	res, err := client.Get(url)
	if err != nil {
		panic(err)
	}

	countries := []Country{}
	err = res.Bind(&countries)
	if err != nil {
		panic(err)
	}

	// callingCodes := []CallingCode{}
	for _, country := range countries {
		// callingCodes = append(callingCodes)
		pretty.Print(country)
	}
}

func main() {

	var (
		url = "https://restcountries.com/v3.1/all?fields=name,flags,idd"
	)

	http.Get(url)
}
