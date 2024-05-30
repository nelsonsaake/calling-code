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

	toCallingCodes := func(country models.Country, suffix string) models.CallingCode {
		return models.CallingCode{
			Flag:        country.Flags.Png,
			CountryCode: country.CountryCode,
			CallingCode: country.Idd.Root + suffix,
		}
	}

	for _, country := range countries {

		if len(country.Idd.Suffixes) == 0 {
			callingCodes = append(callingCodes, toCallingCodes(country, ""))
			continue
		}

		for _, suf := range country.Idd.Suffixes {
			callingCodes = append(callingCodes, toCallingCodes(country, suf))
		}
	}

	err = ufs.WriteFile("calling_codes.json", pretty.JSON(callingCodes))
	if err != nil {
		panic(err)
	}
}
