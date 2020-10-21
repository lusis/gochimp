package api

// CartInfoResponse ...
// https://mailchimp.com/developer/api/marketing/ecommerce-carts/get-cart-info/
type CartInfoResponse struct {
	ID           string                     `json:"id"`
	Customer     CustomerInfoResponse       `json:"customer"`
	CampaignID   string                     `json:"campaign_id"`
	CheckoutURL  string                     `json:"checkout_url"`
	CurrencyCode string                     `json:"currency_code"`
	OrderTotal   float32                    `json:"order_total"`
	TaxTotal     float32                    `json:"tax_total"`
	Lines        []CartLineItemInfoResponse `json:"lines"`
	CreatedAt    Time                       `json:"created_at"`
	UpdatedAt    Time                       `json:"updated_at"`
	CommonResponse
}

// ListCartsResponse ...
// https://mailchimp.com/developer/api/marketing/ecommerce-carts/list-carts/
type ListCartsResponse struct {
	StoreID string             `json:"id"`
	Carts   []CartInfoResponse `json:"carts"`
	CommonCollectionResponse
}

// AddCartRequest ...
// https://mailchimp.com/developer/api/marketing/ecommerce-carts/add-cart/
type AddCartRequest struct {
	ID           string                   `json:"id"`
	Customer     AddCustomerRequest       `json:"customer"`
	CurrencyCode string                   `json:"currency_code"`
	OrderTotal   float32                  `json:"order_total"`
	CampaignID   string                   `json:"campaign_id,omitempty"`
	CheckoutURL  string                   `json:"checkout_url,omitempty"`
	TaxTotal     float32                  `json:"tax_total,omitempty"`
	Lines        []AddCartLineItemRequest `json:"lines"`
}

// AddCartResponse ...
type AddCartResponse CartInfoResponse

// UpdateCartRequest ...
// https://mailchimp.com/developer/api/marketing/ecommerce-carts/update-cart/
type UpdateCartRequest struct {
	Customer     UpdateCustomerRequest       `json:"customer,omitempty"`
	CurrencyCode string                      `json:"currency_code,omitempty"`
	OrderTotal   float32                     `json:"order_total,omitempty"`
	CampaignID   string                      `json:"campaign_id,omitempty"`
	CheckoutURL  string                      `json:"checkout_url,omitempty"`
	TaxTotal     float32                     `json:"tax_total,omitempty"`
	Lines        []UpdateCartLineItemRequest `json:"lines,omitempty"`
}

// UpdateCartResponse ...
type UpdateCartResponse CartInfoResponse

// DeleteCartResponse ...
type DeleteCartResponse EmptyResponse
