package api

// CompanyAddressResponse ...
/*
these are address structs that contain additional company data
*/
type CompanyAddressResponse struct {
	Name      string  `json:"name"`
	Company   string  `json:"company"`
	Phone     string  `json:"phone"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	CommonAddressResponse
}

// OrderInfoResponse ...
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
  "store_id": "string",
  "campaign_id": "839488a60b",
  "landing_site": "http://www.example.com?source=abc",
  "financial_status": "string",
  "fulfillment_status": "string",
  "currency_code": "string",
  "order_total": 0,
  "order_url": "string",
  "discount_total": 0,
  "tax_total": 0,
  "shipping_total": 0,
  "tracking_code": "prec",
  "processed_at_foreign": "2015-07-15 19:28:00",
  "cancelled_at_foreign": "2015-07-15 19:28:00",
  "updated_at_foreign": "2015-07-15 19:28:00",
  "shipping_address": {
    "name": "Freddie Chimpenheimer",
    "address1": "675 Ponce de Leon Ave NE",
    "address2": "Suite 5000",
    "city": "Atlanta",
    "province": "Georgia",
    "province_code": "GA",
    "postal_code": "30308",
    "country": "United States",
    "country_code": "US",
    "longitude": -75.68903,
    "latitude": 45.427408,
    "phone": "8675309",
    "company": "string"
  },
  "billing_address": {
    "name": "Freddie Chimpenheimer",
    "address1": "675 Ponce de Leon Ave NE",
    "address2": "Suite 5000",
    "city": "Atlanta",
    "province": "Georgia",
    "province_code": "GA",
    "postal_code": "30308",
    "country": "United States",
    "country_code": "US",
    "longitude": -75.68903,
    "latitude": 45.427408,
    "phone": "8675309",
    "company": "string"
  },
  "promos": [
    {
      "code": "string",
      "amount_discounted": 0,
      "type": "fixed"
    }
  ],
  "lines": [
    {
      "id": "string",
      "product_id": "string",
      "product_title": "string",
      "product_variant_id": "string",
      "product_variant_title": "string",
      "image_url": "string",
      "quantity": 0,
      "price": 0,
      "discount": 0,
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
  "outreach": {
    "id": "839488a60b",
    "name": "Freddie's Jokes",
    "type": "regular",
    "published_time": "2017-06-06T13:56:12+00:00"
  },
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
type OrderInfoResponse struct {
	ID                 string                      `json:"id"`
	Customer           CustomerInfoResponse        `json:"customer"`
	StoreID            string                      `json:"store_id"`
	CampaignID         string                      `json:"campaign_id"`
	LandingSite        string                      `json:"landing_site"`
	FinancialStatus    string                      `json:"financial_status"`
	FullfillmentStatus string                      `json:"fullfillment_status"`
	CurrencyCode       string                      `json:"currency_code"`
	OrderTotal         float32                     `json:"order_total"`
	OrderURL           string                      `json:"order_url"`
	DiscountTotal      float32                     `json:"discount_total"`
	TaxTotal           float32                     `json:"tax_total"`
	ShippingTotal      float32                     `json:"shipping_total"`
	ProcessedAtForeign Time                        `json:"proccessed_at_foreign"`
	CancelledAtForeign Time                        `json:"cancelled_at_foreign"`
	UpdatedAtForeign   Time                        `json:"updated_at_foreign"`
	ShippingAddress    CompanyAddressResponse      `json:"shipping_address"`
	BillingAddress     CompanyAddressResponse      `json:"billing_address"`
	Outreach           OrderOutreachResponse       `json:"outreach"`
	Promos             []OrderPromoItemResponse    `json:"promos"`
	Lines              []OrderLineItemInfoResponse `json:"lines"`
	CommonResponse
}

// OrderPromoItemResponse ...
type OrderPromoItemResponse struct {
	Code             string  `json:"code"`
	AmountDiscounted float32 `json:"amount_discounted"`
	Type             string  `json:"type"`
}

// OrderOutreachResponse ...
type OrderOutreachResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	PublishedAt Time   `json:"published_at"`
}
