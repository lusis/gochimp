package api

// ProductInfoResponse ...
// https://mailchimp.com/developer/api/marketing/ecommerce-products/get-product-info/
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
type ListProductsResponse struct {
	StoreID  string                `json:"store_id"`
	Products []ProductInfoResponse `json:"products"`
	CommonCollectionResponse
}

// AddProductRequest ...
// https://mailchimp.com/developer/api/marketing/ecommerce-products/add-product/
type AddProductRequest struct {
	ID                 string                     `json:"id"`
	Title              string                     `json:"title"`
	Variants           []AddProductVariantRequest `json:"variants"`
	Handle             string                     `json:"handle,omitempty"`
	URL                string                     `json:"url,omitempty"`
	Description        string                     `json:"description,omitempty"`
	Type               string                     `json:"type,omitempty"`
	Vendor             string                     `json:"vendor,omitempty"`
	ImageURL           string                     `json:"image_url,omitempty"`
	Images             []AddProductImageRequest   `json:"images,omitempty"`
	PublishedAtForeign Time                       `json:"published_at_foreign,omitempty"`
}

// AddProductResponse ...
type AddProductResponse ProductInfoResponse

// UpdateProductRequest ...
// https://mailchimp.com/developer/api/marketing/ecommerce-products/update-product/
type UpdateProductRequest struct {
	Title              string                        `json:"title"`
	Variants           []UpdateProductVariantRequest `json:"variants"`
	Handle             string                        `json:"handle,omitempty"`
	URL                string                        `json:"url,omitempty"`
	Description        string                        `json:"description,omitempty"`
	Type               string                        `json:"type,omitempty"`
	Vendor             string                        `json:"vendor,omitempty"`
	ImageURL           string                        `json:"image_url,omitempty"`
	Images             []UpdateProductImageRequest   `json:"images,omitempty"`
	PublishedAtForeign Time                          `json:"published_at_foreign,omitempty"`
}

// UpdateProductResponse ...
type UpdateProductResponse ProductInfoResponse

// DeleteProductResponse ...
type DeleteProductResponse EmptyResponse
