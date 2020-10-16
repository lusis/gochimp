package api

// CustomerInfoResponse ...
/*
{
  "id": "string",
  "email_address": "string",
  "opt_in_status": true,
  "company": "string",
  "first_name": "string",
  "last_name": "string",
  "orders_count": 4,
  "total_spent": 100,
  "address": {
    "address1": "675 Ponce de Leon Ave NE",
    "address2": "Suite 5000",
    "city": "Atlanta",
    "province": "Georgia",
    "province_code": "GA",
    "postal_code": "30308",
    "country": "United States",
    "country_code": "US"
  },
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
type CustomerInfoResponse struct {
	ID           string                  `json:"id"`
	EmailAddress string                  `json:"email_address"`
	OptInStatus  bool                    `json:"opt_in_status"`
	Company      string                  `json:"company"`
	FirstName    string                  `json:"first_name"`
	LastName     string                  `json:"last_name"`
	OrdersCount  int32                   `json:"orders_count"`
	TotalSpent   float32                 `json:"total_spent"`
	Address      CustomerAddressResponse `json:"address"`
	CreatedAt    Time                    `json:"created_at"`
	UpdatedAt    Time                    `json:"updated_at"`
	CommonResponse
}

// CustomerAddressResponse ...
type CustomerAddressResponse struct {
	CommonAddressResponse
}

// ListCustomersResponse ...
/*
{
  "store_id": "string",
  "customers": [
    {
      "id": "string",
      "email_address": "string",
      "opt_in_status": true,
      "company": "string",
      "first_name": "string",
      "last_name": "string",
      "orders_count": 4,
      "total_spent": 100,
      "address": {
        "address1": "675 Ponce de Leon Ave NE",
        "address2": "Suite 5000",
        "city": "Atlanta",
        "province": "Georgia",
        "province_code": "GA",
        "postal_code": "30308",
        "country": "United States",
        "country_code": "US"
      },
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
  ],
  "total_items": 0,
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
type ListCustomersResponse struct {
	StoreID    string                 `json:"store_id"`
	Customers  []CustomerInfoResponse `json:"customers"`
	TotalItems int32                  `json:"total_items"`
	CommonResponse
}
