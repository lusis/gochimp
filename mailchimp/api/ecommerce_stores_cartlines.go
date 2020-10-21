package api

// CartLineItemInfoResponse ...
// https://mailchimp.com/developer/api/marketing/ecommerce-cart-lines/get-cart-line-item/
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
// https://mailchimp.com/developer/api/marketing/ecommerce-cart-lines/list-cart-line-items/
type ListCartLineItemsResponse struct {
	StoreID string                     `json:"store_id"`
	CartID  string                     `json:"cart_id"`
	Lines   []CartLineItemInfoResponse `json:"lines"`
	CommonCollectionResponse
}

// AddCartLineItemRequest ...
// https://mailchimp.com/developer/api/marketing/ecommerce-cart-lines/add-cart-line-item/
type AddCartLineItemRequest struct {
	ID               string  `json:"id"`
	ProductID        string  `json:"product_id"`
	ProductVariantID string  `json:"product_variant_id"`
	Quantity         int32   `json:"quantity"`
	Price            float32 `json:"price"`
}

// AddCartLineItemResponse ...
type AddCartLineItemResponse CartLineItemInfoResponse

// UpdateCartLineItemRequest ...
// https://mailchimp.com/developer/api/marketing/ecommerce-cart-lines/update-cart-line-item/
type UpdateCartLineItemRequest struct {
	ProductID        string  `json:"product_id,omitempty"`
	ProductVariantID string  `json:"product_variant_id,omitempty"`
	Quantity         int32   `json:"quantity,omitempty"`
	Price            float32 `json:"price,omitempty"`
}

// UpdateCartLineItemResponse ...
type UpdateCartLineItemResponse CartLineItemInfoResponse

// DeleteCartLineItemResponse ...
type DeleteCartLineItemResponse EmptyResponse
