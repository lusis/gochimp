package api

// OrderLineItemInfoResponse ...
/*
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
*/
type OrderLineItemInfoResponse struct {
	ID                  string  `json:"id"`
	ProductID           string  `json:"product_id"`
	ProductTitle        string  `json:"product_title"`
	ProductVariantID    string  `json:"product_variant_id"`
	ProductVariantTitle string  `json:"product_variant_title"`
	ImageURL            string  `json:"image_url"`
	Quantity            int32   `json:"quantity"`
	Price               float32 `json:"price"`
	Discount            float32 `json:"discount"`
	CommonResponse
}
