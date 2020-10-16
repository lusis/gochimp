package api

// CommonResponse ...
type CommonResponse struct {
	Links []LinkResponse `json:"_links"`
}

// CommonAddressResponse ...
/*
these are fields common to all address structs
some struct have additional fields like lat/long
*/
type CommonAddressResponse struct {
	AddressOne   string `json:"address1"`
	AddressTwo   string `json:"address2"`
	City         string `json:"city"`
	Province     string `json:"province"`
	ProvinceCode string `json:"province_code"`
	Country      string `json:"country"`
	CountryCode  string `json:"country_code"`
}
