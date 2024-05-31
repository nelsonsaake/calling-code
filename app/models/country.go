package models

type Country struct {
	CountryCode string `json:"cca2"`
	Name        struct {
		Common string `json:"common"`
	} `json:"name"`
	Flags struct {
		Png string `json:"png"`
	} `json:"flags"`
	Idd struct {
		Root     string   `json:"root"`
		Suffixes []string `json:"suffixes"`
	} `json:"idd"`
}
