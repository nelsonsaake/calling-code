package main

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	_ "embed"

	"github.com/nelsonsaake/calling-code/app/models"
	"github.com/nelsonsaake/go-ns/dld"
	"github.com/nelsonsaake/go-ns/pretty"
	"github.com/nelsonsaake/go-ns/str"
	"github.com/nelsonsaake/go-ns/ufs"
)

var (
	//go:embed calling_codes.json
	callingCodeJson string

	callingCodeBytes = []byte(callingCodeJson)

	callingCodes = []models.CallingCode{}

	flagStorage = "storage/flags"
)

func main() {

	err := json.Unmarshal(callingCodeBytes, &callingCodes)
	if err != nil {
		panic(err)
	}

	for _, callingCode := range callingCodes {

		fmt.Printf("downloading %v\n", callingCode.Flag)

		if str.IsEmpty(callingCode.Flag) {
			pretty.Print(callingCode)
			continue
		}

		src, err := dld.DownloadFile(callingCode.Flag, flagStorage)
		if err != nil {
			panic(fmt.Errorf("error downloading file: %v", err))
		}

		dst := filepath.Join(flagStorage, callingCode.CountryCode+filepath.Ext(src))
		dst = strings.ToLower(dst)

		err = ufs.Rename(src, dst)
		if err != err {
			panic(err)
		}
	}
}
