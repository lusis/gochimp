package api

// EcommerceListProductVariantsResponse ...
// https://mailchimp.com/developer/api/marketing/ecommerce-product-variants/list-product-variants/
/*
{
  "store_id": "string",
  "product_id": "string",
  "variants": [
    {
      "id": "string",
      "title": "Cat Hat",
      "url": "string",
      "sku": "string",
      "price": 0,
      "inventory_quantity": 0,
      "image_url": "string",
      "backorders": "string",
      "visibility": "string",
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
type EcommerceListProductVariantsResponse struct {
	StoreID    string                       `json:"store_id"`
	ProductID  string                       `json:"product_id"`
	Variants   []ProductVariantInfoResponse `json:"variants"`
	TotalItems int32                        `json:"total_items"`
	CommonResponse
}

// ProductVariantInfoResponse ...
// https://mailchimp.com/developer/api/marketing/ecommerce-product-variants/get-product-variant-info/
/*
{
  "id": "string",
  "title": "Cat Hat",
  "url": "string",
  "sku": "string",
  "price": 0,
  "inventory_quantity": 0,
  "image_url": "string",
  "backorders": "string",
  "visibility": "string",
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
type ProductVariantInfoResponse struct {
	ID                string  `json:"id"`
	Title             string  `json:"title"`
	URL               string  `json:"url"`
	SKU               string  `json:"sku"`
	Price             float32 `json:"price"`
	InventoryQuantity int32   `json:"inventory_quantity"`
	ImageURL          string  `json:"image_url"`
	Backorders        string  `json:"backorders"`
	Visibility        string  `json:"visibility"`
	CreatedAt         Time    `json:"created_at"`
	UpdatedAt         Time    `json:"updated_at"`
	CommonResponse
}

// AddOrUpdateProductVariantRequest ...
// https://mailchimp.com/developer/api/marketing/ecommerce-product-variants/add-product-variant/
/*
{
	"id":"",
	"title":"",
	"url":"",
	"sku":"",
	"price":0,
	"inventory_quantity":0,
	"image_url":"",
	"backorders":"",
	"visibility":""
}
*/
type AddOrUpdateProductVariantRequest struct {
	ID                string  `json:"id"`
	Title             string  `json:"title"`
	URL               string  `json:"url,omitempty"`
	SKU               string  `json:"sku,omitempty"`
	Price             float32 `json:"price,omitempty"`
	InventoryQuantity int32   `json:"inventory_quantity,omitempty"`
	ImageURL          string  `json:"image_url,omitempty"`
	Backorders        string  `json:"backorders,omitempty"`
	Visibility        string  `json:"visibility,omitempty"`
}
