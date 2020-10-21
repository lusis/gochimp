package api

// OrderInfoResponse ...
// https://mailchimp.com/developer/api/marketing/ecommerce-orders/get-order-info/
type OrderInfoResponse struct {
	ID                 string                       `json:"id"`
	Customer           CustomerInfoResponse         `json:"customer"`
	StoreID            string                       `json:"store_id"`
	CampaignID         string                       `json:"campaign_id"`
	LandingSite        string                       `json:"landing_site"`
	FinancialStatus    string                       `json:"financial_status"`
	FullfillmentStatus string                       `json:"fullfillment_status"`
	CurrencyCode       string                       `json:"currency_code"`
	OrderTotal         float32                      `json:"order_total"`
	OrderURL           string                       `json:"order_url"`
	DiscountTotal      float32                      `json:"discount_total"`
	TaxTotal           float32                      `json:"tax_total"`
	ShippingTotal      float32                      `json:"shipping_total"`
	ProcessedAtForeign Time                         `json:"proccessed_at_foreign"`
	CancelledAtForeign Time                         `json:"cancelled_at_foreign"`
	UpdatedAtForeign   Time                         `json:"updated_at_foreign"`
	ShippingAddress    CommonCompanyAddressResponse `json:"shipping_address"`
	BillingAddress     CommonCompanyAddressResponse `json:"billing_address"`
	Outreach           OrderOutreachResponse        `json:"outreach"`
	Promos             []OrderPromoItemResponse     `json:"promos"`
	Lines              []OrderLineItemInfoResponse  `json:"lines"`
	CommonResponse
}

// OrderPromoItemResponse ...
type OrderPromoItemResponse struct {
	Code             string  `json:"code"`
	AmountDiscounted float32 `json:"amount_discounted"`
	Type             string  `json:"type"`
}

// OrderOutreachResponse ...
type OrderOutreachResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	PublishedAt Time   `json:"published_at"`
}

// ListOrdersResponse ...
// https://mailchimp.com/developer/api/marketing/ecommerce-orders/list-orders/
type ListOrdersResponse struct {
	StoreID string              `json:"store_id"`
	Orders  []OrderInfoResponse `json:"orders"`
	CommonCollectionResponse
}

// AddOrderRequest ...
// https://mailchimp.com/developer/api/marketing/ecommerce-orders/add-order/
type AddOrderRequest struct{}

// AddOrderResponse ...
type AddOrderResponse OrderInfoResponse

// UpdateOrderRequest ...
// https://mailchimp.com/developer/api/marketing/ecommerce-orders/update-order/
type UpdateOrderRequest struct{}

// UpdateOrderResponse ...
type UpdateOrderResponse OrderInfoResponse

// DeleteOrderResponse ...
type DeleteOrderResponse EmptyResponse
