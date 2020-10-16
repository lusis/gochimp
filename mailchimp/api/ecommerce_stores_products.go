package api

// ProductInfoResponse ...
/*
{
  "id": "string",
  "currency_code": "string",
  "title": "Cat Hat",
  "handle": "cat-hat",
  "url": "string",
  "description": "This is a cat hat.",
  "type": "Accessories",
  "vendor": "string",
  "image_url": "string",
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
  "published_at_foreign": "2015-07-15 19:28:00",
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
type ProductInfoResponse struct {
	ID                 string                       `json:"id"`
	CurrencyCode       string                       `json:"currency_code"`
	Title              string                       `json:"title"`
	Handle             string                       `json:"handle"`
	URL                string                       `json:"url"`
	Description        string                       `json:"description"`
	Type               string                       `json:"type"`
	Vendor             string                       `json:"vendor"`
	ImageURL           string                       `json:"image_url"`
	Variants           []ProductVariantInfoResponse `json:"variants"`
	Images             []ProductImageInfoResponse   `json:"images"`
	PublishedAtForeign Time                         `json:"published_at_foreign"`
	CommonResponse
}

// ListProductsResponse ...
// https://mailchimp.com/developer/api/marketing/ecommerce-products/list-product/
/*
{
  "store_id": "string",
  "products": [
    {
      "id": "string",
      "currency_code": "string",
      "title": "Cat Hat",
      "handle": "cat-hat",
      "url": "string",
      "description": "This is a cat hat.",
      "type": "Accessories",
      "vendor": "string",
      "image_url": "string",
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
      "published_at_foreign": "2015-07-15 19:28:00",
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
type ListProductsResponse struct {
	StoreID    string                `json:"store_id"`
	Products   []ProductInfoResponse `json:"products"`
	TotalItems int32                 `json:"total_items"`
	CommonResponse
}

// AddProductRequest ...
// https://mailchimp.com/developer/api/marketing/ecommerce-products/add-product/
/*
{
	"id":"",
	"title":"",
	"handle":"",
	"url":"",
	"description":"",
	"type":"",
	"vendor":"",
	"image_url":"",
	"variants":[],
	"images":[],
	"published_at_foreign":""
}
*/
type AddProductRequest struct {
	ID                 string                             `json:"id"`
	Title              string                             `json:"title"`
	Variants           []AddOrUpdateProductVariantRequest `json:"variants"`
	Handle             string                             `json:"handle,omitempty"`
	URL                string                             `json:"url,omitempty"`
	Description        string                             `json:"description,omitempty"`
	Type               string                             `json:"type,omitempty"`
	Vendor             string                             `json:"vendor,omitempty"`
	ImageURL           string                             `json:"image_url,omitempty"`
	Images             []AddProductImageRequest           `json:"images,omitempty"`
	PublishedAtForeign Time                               `json:"published_at_foreign,omitempty"`
}

// UpdateProductRequest ...
// https://mailchimp.com/developer/api/marketing/ecommerce-products/update-product/
type UpdateProductRequest struct {
	Title              string                             `json:"title"`
	Variants           []AddOrUpdateProductVariantRequest `json:"variants"`
	Handle             string                             `json:"handle,omitempty"`
	URL                string                             `json:"url,omitempty"`
	Description        string                             `json:"description,omitempty"`
	Type               string                             `json:"type,omitempty"`
	Vendor             string                             `json:"vendor,omitempty"`
	ImageURL           string                             `json:"image_url,omitempty"`
	Images             []AddProductImageRequest           `json:"images,omitempty"`
	PublishedAtForeign Time                               `json:"published_at_foreign,omitempty"`
}
