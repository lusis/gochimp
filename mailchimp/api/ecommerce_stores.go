package api

// EcommerceStoreInfoResponse ...
/*
{
  "id": "example_store",
  "list_id": "1a2df69511",
  "name": "Freddie's Cat Hat Emporium",
  "platform": "string",
  "domain": "example.com",
  "is_syncing": true,
  "email_address": "freddie@mailchimp.com",
  "currency_code": "USD",
  "money_format": "$",
  "primary_locale": "fr",
  "timezone": "Eastern",
  "phone": "404-444-4444",
  "address": {
    "address1": "675 Ponce de Leon Ave NE",
    "address2": "Suite 5000",
    "city": "Atlanta",
    "province": "Georgia",
    "province_code": "GA",
    "postal_code": "30308",
    "country": "United States",
    "country_code": "US",
    "longitude": -75.68903,
    "latitude": 45.427408
  },
  "connected_site": {
    "site_foreign_id": "a180c384d7db88b if created in-app, MC001 if created via API",
    "site_script": {
      "url": "https://chimpstatic.com/mcjs-connected/js/users/{user-hash}/{site-hash}.js",
      "fragment": "<script id=\"mcjs\">!function(c,h,i,m,p){m=c.createElement(h),p=c.getElementsByTagName(h)[0],m.async=1,m.src=i,p.parentNode.insertBefore(m,p)}(document,\"script\",\"https://chimpstatic.com/mcjs-connected/js/users/{user-hash}/{site-hash}.js\");</script>"
    }
  },
  "automations": {
    "abandoned_cart": {
      "is_supported": false,
      "id": "355a72bfc3",
      "status": "sending"
    },
    "abandoned_browse": {
      "is_supported": false,
      "id": "355a72bfc3",
      "status": "sending"
    }
  },
  "list_is_active": true,
  "created_at": "2015-07-15 19:28:00",
  "updated_at": "2015-07-15 19:28:00",
  "_links": [
    {
      "rel": "string",
      "href": "string",
      "method": "GET",
      "targetSchema": "string",
      "schema": "string"
    }
  ]
}
*/
type EcommerceStoreInfoResponse struct {
	ID            string                                `json:"id"`
	ListID        string                                `json:"list_id"`
	Name          string                                `json:"name"`
	Platform      string                                `json:"platform"`
	Domain        string                                `json:"domain"`
	IsSyncing     bool                                  `json:"is_syncing"`
	EmailAddress  string                                `json:"email_address"`
	CurrencyCode  string                                `json:"currency_code"`
	MoneyFormat   string                                `json:"money_format"`
	PrimaryLocale string                                `json:"primary_locale"`
	Timezone      string                                `json:"timezone"`
	Phone         string                                `json:"phone"`
	Address       EcommerceAddressDetailRequestResponse `json:"address"`
	ConnectedSite EcommerceConnectedSiteResponse        `json:"connected_site"`
	Automations   EcommerceAutomationsResponse          `json:"automations"`
	ListIsActive  bool                                  `json:"list_is_active"`
	CreatedAt     Time                                  `json:"created_at"`
	UpdatedAt     Time                                  `json:"updated_at"`
	CommonResponse
}

// EcommerceAddressDetailRequestResponse ...
type EcommerceAddressDetailRequestResponse struct {
	AddressOne   string  `json:"address1,omitempty"`
	AddressTwo   string  `json:"address2,omitempty"`
	City         string  `json:"city,omitempty"`
	Province     string  `json:"province,omitempty"`
	ProvinceCode string  `json:"province_code,omitempty"`
	PostalCode   string  `json:"postal_code,omitempty"`
	Country      string  `json:"country,omitempty"`
	CountryCode  string  `json:"country_code,omitempty"`
	Longitude    float64 `json:"longitude,omitempty"`
	Latitude     float64 `json:"latitude,omitempty"`
}

// EcommerceConnectedSiteResponse ...
type EcommerceConnectedSiteResponse struct {
	SiteForeignID string `json:"site_foreign_id"`
	SiteScript    struct {
		URL      string `json:"url"`
		Fragment string `json:"fragment"`
	} `json:"site_script"`
}

// EcommerceAutomationsResponse ...
type EcommerceAutomationsResponse struct {
	AbandonedCart   EcommerceAutomationsDetailResponse `json:"abandoned_cart"`
	AbandonedBrowse EcommerceAutomationsDetailResponse `json:"abandoned_browse"`
}

// EcommerceAutomationsDetailResponse ...
type EcommerceAutomationsDetailResponse struct {
	IsSupported bool   `json:"is_supported"`
	ID          string `json:"id"`
	Status      string `json:"status"`
}

// EcommerceAddStoreRequest ...
/*
{
	"id":"",
	"list_id":"",
	"name":"",
	"platform":"",
	"domain":"",
	"is_syncing":false,
	"email_address":"",
	"currency_code":"",
	"money_format":"",
	"primary_locale":"",
	"timezone":"",
	"phone":"",
	"address":{
		"address1":"",
		"address2":"",
		"city":"",
		"province":"",
		"province_code":"",
		"postal_code":"",
		"country":"",
		"country_code":"",
		"longitude":0,
		"latitude":0
	}
}
*/
type EcommerceAddStoreRequest struct {
	// required fields
	ID           string `json:"id"`
	ListID       string `json:"list_id"`
	Name         string `json:"name"`
	CurrencyCode string `json:"currency_code"`
	// optional fields
	Platform      string                                `json:"platform,omitempty"`
	Domain        string                                `json:"domain,omitempty"`
	IsSyncing     bool                                  `json:"is_syncing,omitempty"`
	EmailAddress  string                                `json:"email_address,omitempty"`
	MoneyFormat   string                                `json:"money_format,omitempty"`
	PrimaryLocale string                                `json:"primary_locale,omitempty"`
	Timezone      string                                `json:"timezone,omitempty"`
	Phone         string                                `json:"phone,omitempty"`
	Address       EcommerceAddressDetailRequestResponse `json:"address,omitempty"`
}

// EcommerceUpdateStoreRequest ...
/*
{
	"name":"",
	"platform":"",
	"domain":"",
	"is_syncing":false,
	"email_address":"",
	"currency_code":"",
	"money_format":"",
	"primary_locale":"",
	"timezone":"",
	"phone":"",
	"address":{
		"address1":"",
		"address2":"",
		"city":"",
		"province":"",
		"province_code":"",
		"postal_code":"",
		"country":"",
		"country_code":"",
		"longitude":0,
		"latitude":0
	}
}
*/
type EcommerceUpdateStoreRequest struct {
	Name          string                                `json:"name,omitempty"`
	CurrencyCode  string                                `json:"currency_code,omitempty"`
	Platform      string                                `json:"platform,omitempty"`
	Domain        string                                `json:"domain,omitempty"`
	IsSyncing     bool                                  `json:"is_syncing,omitempty"`
	EmailAddress  string                                `json:"email_address,omitempty"`
	MoneyFormat   string                                `json:"money_format,omitempty"`
	PrimaryLocale string                                `json:"primary_locale,omitempty"`
	Timezone      string                                `json:"timezone,omitempty"`
	Phone         string                                `json:"phone,omitempty"`
	Address       EcommerceAddressDetailRequestResponse `json:"address,omitempty"`
}

// EcommerceListStoresResponse ...
/*
	https://mailchimp.com/developer/api/marketing/ecommerce-stores/list-stores/
*/
type EcommerceListStoresResponse struct {
	Stores []EcommerceStoreInfoResponse `json:"stores"`
	CommonResponse
}
