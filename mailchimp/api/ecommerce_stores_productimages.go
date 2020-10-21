package api

// ProductImageInfoResponse ...
// https://mailchimp.com/developer/api/marketing/ecommerce-product-images/get-product-image-info/
type ProductImageInfoResponse struct {
	ID         string   `json:"id"`
	URL        string   `json:"url"`
	VariantIDs []string `json:"variant_ids"`
	CommonResponse
}

// ListProductImageInfoResponse ...
// https://mailchimp.com/developer/api/marketing/ecommerce-product-images/list-product-images/
type ListProductImageInfoResponse struct {
	StoreID   string                     `json:"store_id"`
	ProductID string                     `json:"product_id"`
	Images    []ProductImageInfoResponse `json:"images"`
	CommonCollectionResponse
}

// AddProductImageRequest ...
// https://mailchimp.com/developer/api/marketing/ecommerce-product-images/add-product-image/
type AddProductImageRequest struct {
	ID         string   `json:"id"`
	URL        string   `json:"url"`
	VariantIDs []string `json:"variant_ids,omitempty"`
}

// AddProductImageResponse ...
type AddProductImageResponse ProductImageInfoResponse

// UpdateProductImageRequest ...
// https://mailchimp.com/developer/api/marketing/ecommerce-product-images/update-product-image/
type UpdateProductImageRequest struct {
	ID         string   `json:"id,omitempty"`
	URL        string   `json:"url,omitempty"`
	VariantIDs []string `json:"variant_ids,omitempty"`
}

// UpdateProductImageResponse ...
type UpdateProductImageResponse ProductImageInfoResponse

// DeleteProductImageResponse ...
type DeleteProductImageResponse EmptyResponse
