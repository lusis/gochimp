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

// AddCustomerRequest ...
// https://mailchimp.com/developer/api/marketing/ecommerce-customers/add-customer/
/*
{
  "id":"",
  "email_address":"",
  "opt_in_status":false,
  "company":"",
  "first_name":"",
  "last_name":"",
  "address":{
    "address1":"",
    "address2":"",
    "city":"",
    "province":"",
    "province_code":"",
    "postal_code":"",
    "country":"",
    "country_code":""
  }
}
*/
type AddCustomerRequest struct {
	ID           string               `json:"id"`
	EmailAddress string               `json:"email_address"`
	OptInStatus  bool                 `json:"opt_in_status"`
	Company      string               `json:"company,omitempty"`
	FirstName    string               `json:"first_name,omitempty"`
	LastName     string               `json:"last_name,omit_empty"`
	Address      CommonAddressRequest `json:"address,omitempty"`
}

// AddCustomerResponse ...
type AddCustomerResponse CustomerInfoResponse

// AddOrUpdateCustomerRequest ...
// https://mailchimp.com/developer/api/marketing/ecommerce-customers/add-or-update-customer/
type AddOrCustomerRequest struct {
	ID           string               `json:"id"`
	EmailAddress string               `json:"email_address"`
	OptInStatus  bool                 `json:"opt_in_status"`
	Company      string               `json:"company,omitempty"`
	FirstName    string               `json:"first_name,omitempty"`
	LastName     string               `json:"last_name,omit_empty"`
	Address      CommonAddressRequest `json:"address,omitempty"`
}

// AddOrUpdateCustomerResponse ...
type AddOrUpdateCustomerResponse CustomerInfoResponse

// UpdateCustomerRequest ...
// https://mailchimp.com/developer/api/marketing/ecommerce-customers/update-customer/
type UpdateCustomerRequest struct {
	OptInStatus bool                 `json:"opt_in_status,omitempty"`
	Company     string               `json:"company,omitempty"`
	FirstName   string               `json:"first_name,omitempty"`
	LastName    string               `json:"last_name,omit_empty"`
	Address     CommonAddressRequest `json:"address,omitempty"`
}

// UpdateCustomerResponse ...
type UpdateCustomerResponse CustomerInfoResponse

// DeleteCustomerResponse ...
type DeleteCustomerResponse EmptyResponse
