package api

// OrderLineItemInfoResponse ...
// https://mailchimp.com/developer/api/marketing/ecommerce-order-lines/get-order-line-item/
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

// ListOrderLineItemsResponse ...
// https://mailchimp.com/developer/api/marketing/ecommerce-order-lines/list-order-line-items/
type ListOrderLineItemsResponse struct {
	StoreID string                      `json:"store_id"`
	OrderID string                      `json:"order_id"`
	Lines   []OrderLineItemInfoResponse `json:"lines"`
	CommonCollectionResponse
}

// AddOrderLineItemRequest ...
// https://mailchimp.com/developer/api/marketing/ecommerce-order-lines/add-order-line-item/
type AddOrderLineItemRequest struct {
	ID               string  `json:"id"`
	ProductID        string  `json:"product_id"`
	ProductVariantID string  `json:"product_variant_id"`
	Quantity         int32   `json:"quantity"`
	Price            float32 `json:"price"`
	Discount         float32 `json:"discount,omitempty"`
}

// AddOrderLineItemResponse ...
type AddOrderLineItemResponse OrderInfoResponse

// UpdateOrderLineItemRequest ...
// https://mailchimp.com/developer/api/marketing/ecommerce-order-lines/update-order-line-item/
type UpdateOrderLineItemRequest struct {
	ProductID        string  `json:"product_id,omitempty"`
	ProductVariantID string  `json:"product_variant_id,omitempty"`
	Quantity         int32   `json:"quantity,omitempty"`
	Price            float32 `json:"price,omitempty"`
	Discount         float32 `json:"discount,omitempty"`
}

// UpdateOrderLineItemResponse ...
type UpdateOrderLineItemResponse OrderLineItemInfoResponse

// DeleteOrderLineItemResponse ...
type DeleteOrderLineItemResponse EmptyResponse
