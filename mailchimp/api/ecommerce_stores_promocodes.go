package api

// PromoCodeInfoResponse ...
/*
{
  "id": "string",
  "code": "summersale",
  "redemption_url": "A url that applies promo code directly at checkout or a url that points to sale page or store url",
  "usage_count": 0,
  "enabled": "true",
  "created_at_foreign": "2019-08-24T14:15:22Z",
  "updated_at_foreign": "2019-08-24T14:15:22Z",
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
type PromoCodeInfoResponse struct {
	ID               string `json:"id"`
	Code             string `json:"code"`
	RedemptionURL    string `json:"redemption_url"`
	UsageCount       int32  `json:"usage_count"`
	Enabled          bool   `json:"enabled"`
	CreatedAtForeign Time   `json:"created_at_foreign"`
	UpdatedAtForeign Time   `json:"updated_at_foreign"`
	CommonResponse
}

// ListPromoCodesResponse ...
/*
{
  "store_id": "string",
  "promo_codes": [
    {
      "id": "string",
      "code": "summersale",
      "redemption_url": "A url that applies promo code directly at checkout or a url that points to sale page or store url",
      "usage_count": 0,
      "enabled": "true",
      "created_at_foreign": "2019-08-24T14:15:22Z",
      "updated_at_foreign": "2019-08-24T14:15:22Z",
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
type ListPromoCodesResponse struct {
	StoreID    string                  `json:"store_id"`
	PromoCodes []PromoCodeInfoResponse `json:"promo_codes"`
	CommonCollectionResponse
}
