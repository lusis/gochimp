package api

// PromoCodeInfoResponse ...
// https://mailchimp.com/developer/api/marketing/ecommerce-promo-codes/get-promo-code/
type PromoCodeInfoResponse struct {
	ID               string `json:"id"`
	Code             string `json:"code"`
	RedemptionURL    string `json:"redemption_url"`
	UsageCount       int32  `json:"usage_count"`
	Enabled          bool   `json:"enabled"`
	CreatedAtForeign Time   `json:"created_at_foreign"`
	UpdatedAtForeign Time   `json:"updated_at_foreign"`
	CommonResponse
}

// ListPromoCodesResponse ...
// https://mailchimp.com/developer/api/marketing/ecommerce-promo-codes/list-promo-codes/
type ListPromoCodesResponse struct {
	StoreID    string                  `json:"store_id"`
	PromoCodes []PromoCodeInfoResponse `json:"promo_codes"`
	CommonCollectionResponse
}

// AddPromoCodeRequest ...
// https://mailchimp.com/developer/api/marketing/ecommerce-promo-codes/add-promo-code/
// {"id":"","code":"","redemption_url":"","usage_count":0,"enabled":false,"created_at_foreign":"","updated_at_foreign":""}
type AddPromoCodeRequest struct {
	ID               string `json:"id"`
	Code             string `json:"code"`
	RedemptionURL    string `json:"redemption_url"`
	UsageCount       int32  `json:"usage_count,omitempty"`
	Enabled          bool   `json:"enabled,omitempty"`
	CreatedAtForeign Time   `json:"created_at_foreign,omitempty"`
	UpdatedAtForeign Time   `json:"updated_at_foreign,omitempty"`
}

// AddPromoCodeResponse ...
type AddPromoCodeResponse PromoCodeInfoResponse

// UpdatePromoCodeRequest ...
type UpdatePromoCodeRequest struct {
	Code             string `json:"code,omitempty"`
	RedemptionURL    string `json:"redemption_url,omitempty"`
	UsageCount       int32  `json:"usage_count,omitempty"`
	Enabled          bool   `json:"enabled,omitempty"`
	CreatedAtForeign Time   `json:"created_at_foreign"`
	UpdatedAtForeign Time   `json:"updated_at_foreign"`
}

// UpdatePromoCodeResponse ...
type UpdatePromoCodeResponse PromoCodeInfoResponse

// DeletePromoCodeResponse ...
type DeletePromoCodeResponse EmptyResponse
