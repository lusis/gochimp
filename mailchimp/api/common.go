package api

// CommonResponse ...
type CommonResponse struct {
	Links []LinkResponse `json:"_links"`
}

// CommonCollectionResponse ...
type CommonCollectionResponse struct {
	TotalItems int32 `json:"total_items"`
	CommonResponse
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

// CommonAddressRequest ...
type CommonAddressRequest struct {
	AddressOne   string `json:"address1,omitempty"`
	AddressTwo   string `json:"address2,omitempty"`
	City         string `json:"city,omitempty"`
	Province     string `json:"province,omitempty"`
	ProvinceCode string `json:"province_code,omitempty"`
	Country      string `json:"country,omitempty"`
	CountryCode  string `json:"country_code,omitempty"`
}
