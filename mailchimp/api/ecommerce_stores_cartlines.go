package api

// CartLineItemInfoResponse ...
/*
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
*/
type CartLineItemInfoResponse struct {
	ID                  string  `json:"id"`
	ProductID           string  `json:"product_id"`
	ProductTitle        string  `json:"product_title"`
	ProductVariantID    string  `json:"product_variant_id"`
	ProductVariantTitle string  `json:"product_variant_title"`
	Quantity            int32   `json:"quantity"`
	Price               float32 `json:"price"`
	CommonResponse
}

// ListCartLineItemsResponse ...
/*
{
  "store_id": "string",
  "cart_id": "string",
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
type ListCartLineItemsResponse struct {
	StoreID    string                     `json:"store_id"`
	CartID     string                     `json:"cart_id"`
	Lines      []CartLineItemInfoResponse `json:"lines"`
	TotalItems int32                      `json:"total_items"`
	CommonResponse
}
