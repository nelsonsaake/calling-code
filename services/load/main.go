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

		var root = country.Idd.Root
		var suffixes = []string{}

		if len(country.Idd.Suffixes) == 1 {
			root += country.Idd.Suffixes[0]
		} else {
			suffixes = country.Idd.Suffixes
		}

		return models.CallingCode{
			Flag:        country.Flags.Png,
			CountryName: country.Name.Common,
			CountryCode: country.CountryCode,
			Root:        root,
			Suffixes:    suffixes,
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
