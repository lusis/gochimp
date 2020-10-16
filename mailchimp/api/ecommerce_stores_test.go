package api

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecodeListStoresResponse(t *testing.T) {
	data := `
	{
		"stores": [
		  {
			"id": "store_XXXX",
			"list_id": "XXXXXX",
			"name": "Test",
			"platform": "Bigcommerce",
			"domain": "xxxxxxx.mybigcommerce.com",
			"is_syncing": false,
			"email_address": "xxxx@xxxx.com",
			"currency_code": "USD",
			"money_format": "",
			"primary_locale": "",
			"timezone": "Europe/London",
			"phone": "1234567891",
			"address": {
			  "address1": "",
			  "address2": "",
			  "city": "",
			  "province": "",
			  "province_code": "",
			  "postal_code": "",
			  "country": "",
			  "country_code": "",
			  "longitude": 0,
			  "latitude": 0
			},
			"connected_site": {
			  "site_foreign_id": "store_XXXXXXXXX",
			  "site_script": {
				"url": "https://chimpstatic.com/mcjs-connected/js/users/XXXXXXX/XXXXX.js",
				"fragment": "<script id=\"mcjs\">!function(c,h,i,m,p){m=c.createElement(h),p=c.getElementsByTagName(h)[0],m.async=1,m.src=i,p.parentNode.insertBefore(m,p)}(document,\"script\",\"https://chimpstatic.com/mcjs-connected/js/users/XXXXXXXX/XXXXXX.js\");</script>"
			  }
			},
			"automations": {
			  "abandoned_cart": {
				"is_supported": true,
				"id": "XXXXXXX",
				"status": "save"
			  },
			  "abandoned_browse": {
				"is_supported": true,
				"id": "YYYYYYY",
				"status": "sending"
			  }
			},
			"list_is_active": true,
			"created_at": "2018-12-20T14:26:48+00:00",
			"updated_at": "2020-10-12T19:48:13+00:00",
			"_links": [
			  {
				"rel": "self",
				"href": "https://us15.api.mailchimp.com/3.0/ecommerce/stores/store_XXXXXXXXX",
				"method": "GET",
				"targetSchema": "https://us15.api.mailchimp.com/schema/3.0/Definitions/Ecommerce/Stores/Response.json"
			  },
			  {
				"rel": "parent",
				"href": "https://us15.api.mailchimp.com/3.0/ecommerce/stores",
				"method": "GET",
				"targetSchema": "https://us15.api.mailchimp.com/schema/3.0/Definitions/Ecommerce/Stores/CollectionResponse.json"
			  },
			  {
				"rel": "update",
				"href": "https://us15.api.mailchimp.com/3.0/ecommerce/stores/store_XXXXXXXX",
				"method": "PATCH",
				"targetSchema": "https://us15.api.mailchimp.com/schema/3.0/Definitions/Ecommerce/Stores/Response.json",
				"schema": "https://us15.api.mailchimp.com/schema/3.0/Definitions/Ecommerce/Stores/PATCH.json"
			  },
			  {
				"rel": "delete",
				"href": "https://us15.api.mailchimp.com/3.0/ecommerce/stores/store_XXXXXXXXXXX",
				"method": "DELETE"
			  },
			  {
				"rel": "customers",
				"href": "https://us15.api.mailchimp.com/3.0/ecommerce/stores/store_XXXXXXXX/customers",
				"method": "GET",
				"targetSchema": "https://us15.api.mailchimp.com/schema/3.0/Definitions/Ecommerce/Stores/Customers/CollectionResponse.json"
			  },
			  {
				"rel": "products",
				"href": "https://us15.api.mailchimp.com/3.0/ecommerce/stores/store_XXXXXXXXX/products",
				"method": "GET",
				"targetSchema": "https://us15.api.mailchimp.com/schema/3.0/Definitions/Ecommerce/Stores/Products/CollectionResponse.json"
			  },
			  {
				"rel": "promo-rules",
				"href": "https://us15.api.mailchimp.com/3.0/ecommerce/stores/store_XXXXXXXXXX/promo-rules",
				"method": "GET",
				"targetSchema": "https://us15.api.mailchimp.com/schema/3.0/Definitions/Ecommerce/Stores/Promorules/CollectionResponse.json"
			  },
			  {
				"rel": "orders",
				"href": "https://us15.api.mailchimp.com/3.0/ecommerce/stores/store_XXXXXXXX/orders",
				"method": "GET",
				"targetSchema": "https://us15.api.mailchimp.com/schema/3.0/Definitions/Ecommerce/Stores/Orders/CollectionResponse.json"
			  },
			  {
				"rel": "carts",
				"href": "https://us15.api.mailchimp.com/3.0/ecommerce/stores/store_XXXXXXXX/carts",
				"method": "GET",
				"targetSchema": "https://us15.api.mailchimp.com/schema/3.0/Definitions/Ecommerce/Stores/Carts/CollectionResponse.json"
			  }
			]
		  }
		],
		"total_items": 1,
		"_links": [
		  {
			"rel": "self",
			"href": "https://us15.api.mailchimp.com/3.0/ecommerce/stores",
			"method": "GET",
			"targetSchema": "https://us15.api.mailchimp.com/schema/3.0/Definitions/Ecommerce/Stores/CollectionResponse.json"
		  },
		  {
			"rel": "parent",
			"href": "https://us15.api.mailchimp.com/3.0/ecommerce/",
			"method": "GET",
			"targetSchema": "https://us15.api.mailchimp.com/schema/3.0/Definitions/Ecommerce/Namespace.json"
		  },
		  {
			"rel": "create",
			"href": "https://us15.api.mailchimp.com/3.0/ecommerce/stores",
			"method": "POST",
			"targetSchema": "https://us15.api.mailchimp.com/schema/3.0/Definitions/Ecommerce/Stores/Response.json",
			"schema": "https://us15.api.mailchimp.com/schema/3.0/Definitions/Ecommerce/Stores/POST.json"
		  }
		]
	  }
	  
	`
	resp := EcommerceListStoresResponse{}
	err := testStrictDecode([]byte(data), &resp)
	require.NoError(t, err)
}
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
