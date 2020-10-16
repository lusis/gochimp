package api

// PromoRulesInfoResponse ...
/*
{
  "id": "string",
  "title": "50% off Total Order",
  "description": "Save BIG during our summer sale!",
  "starts_at": "2019-08-24T14:15:22Z",
  "ends_at": "string",
  "amount": 0.5,
  "type": "fixed",
  "target": "per_item",
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
type PromoRulesInfoResponse struct {
	ID               string  `json:"id"`
	Title            string  `json:"title"`
	Description      string  `json:"description"`
	StartsAt         Time    `json:"starts_at"`
	EndsAt           Time    `json:"ends_at"`
	Amount           float32 `json:"amount"`
	Type             string  `json:"type"`
	Target           string  `json:"target"`
	Enabled          bool    `json:"enabled"`
	CreatedAtForeign Time    `json:"created_at_foreign"`
	UpdatedAtForeign Time    `json:"updated_at_foreign"`
	CommonResponse
}

// ListPromoRulesResponse ...
/*
{
  "store_id": "string",
  "promo_rules": [
    {
      "id": "string",
      "title": "50% off Total Order",
      "description": "Save BIG during our summer sale!",
      "starts_at": "2019-08-24T14:15:22Z",
      "ends_at": "string",
      "amount": 0.5,
      "type": "fixed",
      "target": "per_item",
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
type ListPromoRulesResponse struct {
	StoreID    string                   `json:"store_id"`
	PromoRules []PromoRulesInfoResponse `json:"promo_rules"`
	CommonCollectionResponse
}
