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
type AddOrderRequest struct {
	ID                 string                         `json:"id"`
	Customer           AddCustomerRequest             `json:"customer"`
	CurrencyCode       string                         `json:"currency_code"`
	OrderTotal         float32                        `json:"order_total"`
	Lines              []AddOrderLineItemRequest      `json:"lines"`
	CampaignID         string                         `json:"campaign_id,omitempty"`
	LandingSite        string                         `json:"landing_site,omitempty"`
	FinancialStatus    string                         `json:"financial_status,omitempty"`
	FullfillmentStatus string                         `json:"fullfillment_status,omitempty"`
	OrderURL           string                         `json:"order_url,omitempty"`
	DiscountTotal      float32                        `json:"discount_total,omitempty"`
	TaxTotal           float32                        `json:"tax_total,omitempty"`
	ShippingTotal      float32                        `json:"shipping_total,omitempty"`
	TrackingCode       string                         `json:"tracking_code,omitempty"`
	ProcessedAtForeign Time                           `json:"processed_at_foreign,omitempty"`
	CancelledAtForeign Time                           `json:"cancelled_at_foreign,omitempty"`
	UpdatedAtForeign   Time                           `json:"updated_at_foreign,omitempty"`
	ShippingAddress    CommonCompanyAddressRequest    `json:"shipping_address,omitempty"`
	BillingAddress     CommonCompanyAddressRequest    `json:"billing_address,omitempty"`
	Promos             []AddOrderPromoCodeItemRequest `json:"promos,omitempty"`
	Outreach           []string                       `json:"outreach,omitempty"`
}

// AddOrderPromoCodeItemRequest ...
type AddOrderPromoCodeItemRequest struct {
	Code             string  `json:"code,omitempty"`
	Type             string  `json:"type,omitempty"`
	AmountDiscounted float32 `json:"amount_discounted,omitempty"`
}

// AddOrderResponse ...
type AddOrderResponse OrderInfoResponse

// UpdateOrderRequest ...
// https://mailchimp.com/developer/api/marketing/ecommerce-orders/update-order/
type UpdateOrderRequest struct {
	Customer           UpdateCustomerRequest             `json:"customer,omitempty"`
	CurrencyCode       string                            `json:"currency_code,omitempty"`
	OrderTotal         float32                           `json:"order_total,omitempty"`
	Lines              []UpdateOrderLineItemRequest      `json:"lines,omitempty"`
	CampaignID         string                            `json:"campaign_id,omitempty"`
	LandingSite        string                            `json:"landing_site,omitempty"`
	FinancialStatus    string                            `json:"financial_status,omitempty"`
	FullfillmentStatus string                            `json:"fullfillment_status,omitempty"`
	OrderURL           string                            `json:"order_url,omitempty"`
	DiscountTotal      float32                           `json:"discount_total,omitempty"`
	TaxTotal           float32                           `json:"tax_total,omitempty"`
	ShippingTotal      float32                           `json:"shipping_total,omitempty"`
	TrackingCode       string                            `json:"tracking_code,omitempty"`
	ProcessedAtForeign Time                              `json:"processed_at_foreign,omitempty"`
	CancelledAtForeign Time                              `json:"cancelled_at_foreign,omitempty"`
	UpdatedAtForeign   Time                              `json:"updated_at_foreign,omitempty"`
	ShippingAddress    CommonCompanyAddressRequest       `json:"shipping_address,omitempty"`
	BillingAddress     CommonCompanyAddressRequest       `json:"billing_address,omitempty"`
	Promos             []UpdateOrderPromoCodeItemRequest `json:"promos,omitempty"`
	Outreach           []string                          `json:"outreach,omitempty"`
}

// UpdateOrderPromoCodeItemRequest ...
type UpdateOrderPromoCodeItemRequest struct {
	Code             string  `json:"code,omitempty"`
	Type             string  `json:"type,omitempty"`
	AmountDiscounted float32 `json:"amount_discounted,omitempty"`
}

// UpdateOrderResponse ...
type UpdateOrderResponse OrderInfoResponse

// DeleteOrderResponse ...
type DeleteOrderResponse EmptyResponse
