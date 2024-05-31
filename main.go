package main

import (
	"encoding/json"

	"github.com/nelsonsaake/calling-code/app/models"
	"github.com/nelsonsaake/go-ns/pretty"
	"github.com/nelsonsaake/go-ns/ufs"
)

var (
	countriesJsonFile = "countries.json"
	countries         = []models.Country{}
	callingCodes      = []models.CallingCode{}
)

func main() {

	bs, err := ufs.ReadFileAsBytes(countriesJsonFile)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bs, &countries)
	if err != nil {
		panic(err)
	}

	toCallingCodes := func(country models.Country) models.CallingCode {
		return models.CallingCode{
			Flag:        country.Flags.Png,
			CountryCode: country.CountryCode,
			Root:        country.Idd.Root,
			Suffixes:    country.Idd.Suffixes,
		}
	}

	for _, country := range countries {

		callingCodes = append(callingCodes, toCallingCodes(country))
	}

	err = ufs.WriteFile("calling_codes.json", pretty.JSON(callingCodes))
	if err != nil {
		panic(err)
	}
}
