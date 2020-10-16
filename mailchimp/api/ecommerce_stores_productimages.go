package api

// ProductImageInfoResponse ...
type ProductImageInfoResponse struct {
	ID         string   `json:"id"`
	URL        string   `json:"url"`
	VariantIDs []string `json:"variant_ids"`
	CommonResponse
}

// ListProductImageInfoResponse ...
// https://mailchimp.com/developer/api/marketing/ecommerce-product-images/list-product-images/
/*
{
  "store_id": "string",
  "product_id": "string",
  "images": [
    {
      "id": "string",
      "url": "string",
      "variant_ids": [
        "string"
      ],
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
type ListProductImageInfoResponse struct {
	StoreID    string                     `json:"store_id"`
	ProductID  string                     `json:"product_id"`
	Images     []ProductImageInfoResponse `json:"images"`
	TotalItems int32                      `json:"total_items"`
	CommonResponse
}

// AddProductImageRequest ...
/*
{"id":"","url":"","variant_ids":[]}
*/
type AddProductImageRequest struct {
	ID         string   `json:"id"`
	URL        string   `json:"url"`
	VariantIDs []string `json:"variant_ids,omitempty"`
}
