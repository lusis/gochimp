package api

// CartInfoResponse ...
/*
{
  "id": "string",
  "customer": {
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
  },
  "campaign_id": "839488a60b",
  "checkout_url": "string",
  "currency_code": "string",
  "order_total": 0,
  "tax_total": 0,
  "lines": [
    {
      "id": "string",
      "product_id": "string",
      "product_title": "string",
      "product_variant_id": "string",
      "product_variant_title": "string",
      "quantity": 0,
      "price": 0,
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
type CartInfoResponse struct {
	ID           string                     `json:"id"`
	Customer     CustomerInfoResponse       `json:"customer"`
	CampaignID   string                     `json:"campaign_id"`
	CheckoutURL  string                     `json:"checkout_url"`
	CurrencyCode string                     `json:"currency_code"`
	OrderTotal   float32                    `json:"order_total"`
	TaxTotal     float32                    `json:"tax_total"`
	Lines        []CartLineItemInfoResponse `json:"lines"`
	CreatedAt    Time                       `json:"created_at"`
	UpdatedAt    Time                       `json:"updated_at"`
	CommonResponse
}

// ListCartsResponse ...
/*
{
  "store_id": "string",
  "carts": [
    {
      "id": "string",
      "customer": {
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
      },
      "campaign_id": "839488a60b",
      "checkout_url": "string",
      "currency_code": "string",
      "order_total": 0,
      "tax_total": 0,
      "lines": [
        {
          "id": "string",
          "product_id": "string",
          "product_title": "string",
          "product_variant_id": "string",
          "product_variant_title": "string",
          "quantity": 0,
          "price": 0,
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
type ListCartsResponse struct {
	StoreID string             `json:"id"`
	Carts   []CartInfoResponse `json:"carts"`
	CommonCollectionResponse
}

// AddCartRequest ...
/*
{
  "id":"",
  "customer":{
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
  },
  "campaign_id":"",
  "checkout_url":"",
  "currency_code":"",
  "order_total":0,
  "tax_total":0,
  "lines":[]
}
*/
type AddCartRequest struct {
	ID           string                   `json:"id"`
	Customer     AddCustomerRequest       `json:"customer"`
	CurrencyCode string                   `json:"currency_code"`
	OrderTotal   float32                  `json:"order_total"`
	CampaignID   string                   `json:"campaign_id,omitempty"`
	CheckoutURL  string                   `json:"checkout_url,omitempty"`
	TaxTotal     float32                  `json:"tax_total,omitempty"`
	Lines        []AddCartLineItemRequest `json:"lines"`
}

// AddCartResponse ...
type AddCartResponse CartInfoResponse

// UpdateCartRequest ...
type UpdateCartRequest struct {
	Customer     UpdateCustomerRequest       `json:"customer,omitempty"`
	CurrencyCode string                      `json:"currency_code,omitempty"`
	OrderTotal   float32                     `json:"order_total,omitempty"`
	CampaignID   string                      `json:"campaign_id,omitempty"`
	CheckoutURL  string                      `json:"checkout_url,omitempty"`
	TaxTotal     float32                     `json:"tax_total,omitempty"`
	Lines        []UpdateCartLineItemRequest `json:"lines,omitempty"`
}

// UpdateCartResponse ...
type UpdateCartResponse CartInfoResponse
