package api

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecodeStoreInfoResponse(t *testing.T) {
	data := `
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
	`
	resp := EcommerceStoreInfoResponse{}
	err := testStrictDecode([]byte(data), &resp)
	require.NoError(t, err)
}
