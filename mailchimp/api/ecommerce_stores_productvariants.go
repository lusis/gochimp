package api

// ListProductVariantsResponse ...
// https://mailchimp.com/developer/api/marketing/ecommerce-product-variants/list-product-variants/
type ListProductVariantsResponse struct {
	StoreID   string                       `json:"store_id"`
	ProductID string                       `json:"product_id"`
	Variants  []ProductVariantInfoResponse `json:"variants"`
	CommonCollectionResponse
}

// ProductVariantInfoResponse ...
// https://mailchimp.com/developer/api/marketing/ecommerce-product-variants/get-product-variant-info/
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

// AddProductVariantRequest ...
// https://mailchimp.com/developer/api/marketing/ecommerce-product-variants/add-product-variant/
type AddProductVariantRequest struct {
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

// AddProductVariantResponse ...
type AddProductVariantResponse ProductVariantInfoResponse

// UpdateProductVariantRequest ...
// https://mailchimp.com/developer/api/marketing/ecommerce-product-variants/update-product-variant/
type UpdateProductVariantRequest struct {
	Title             string  `json:"title,omitempty"`
	URL               string  `json:"url,omitempty"`
	SKU               string  `json:"sku,omitempty"`
	Price             float32 `json:"price,omitempty"`
	InventoryQuantity int32   `json:"inventory_quantity,omitempty"`
	ImageURL          string  `json:"image_url,omitempty"`
	Backorders        string  `json:"backorders,omitempty"`
	Visibility        string  `json:"visibility,omitempty"`
}

// UpdateProductVariantResponse ...
type UpdateProductVariantResponse ProductVariantInfoResponse

// AddOrUpdateProductVariantRequest ...
// https://mailchimp.com/developer/api/marketing/ecommerce-product-variants/add-or-update-product-variant/
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

// AddOrUpdateProductVariantResponse ...
type AddOrUpdateProductVariantResponse ProductVariantInfoResponse

// DeleteProductVariantResponse ...
type DeleteProductVariantResponse EmptyResponse
