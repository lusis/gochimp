package api

// PromoRuleInfoResponse ...
// https://mailchimp.com/developer/api/marketing/ecommerce-promo-rules/get-promo-rule/
type PromoRuleInfoResponse struct {
	ID               string  `json:"id"`
	Title            string  `json:"title"`
	Description      string  `json:"description"`
	StartsAt         Time    `json:"starts_at"`
	EndsAt           Time    `json:"ends_at"`
	Amount           float32 `json:"amount"`
	Type             string  `json:"type"`
	Target           string  `json:"target"`
	Enabled          bool    `json:"enabled"`
	CreatedAtForeign Time    `json:"created_at_foreign"`
	UpdatedAtForeign Time    `json:"updated_at_foreign"`
	CommonResponse
}

// ListPromoRulesResponse ...
// https://mailchimp.com/developer/api/marketing/ecommerce-promo-rules/list-promo-rules/
type ListPromoRulesResponse struct {
	StoreID    string                  `json:"store_id"`
	PromoRules []PromoRuleInfoResponse `json:"promo_rules"`
	CommonCollectionResponse
}

// AddPromoRuleRequest ...
// https://mailchimp.com/developer/api/marketing/ecommerce-promo-rules/add-promo-rule/
type AddPromoRuleRequest struct {
	ID               string  `json:"id"`
	Description      string  `json:"description"`
	Amount           float32 `json:"amount"`
	Type             string  `json:"type"`
	Target           string  `json:"target"`
	Title            string  `json:"title,omitempty"`
	StartsAt         Time    `json:"starts_at,omitempty"`
	EndsAt           Time    `json:"ends_at,omitempty"`
	Enabled          bool    `json:"enabled,omitempty"`
	CreatedAtForeign Time    `json:"created_at_foreign,omitempty"`
	UpdatedAtForeign Time    `json:"updated_at_foreign,omitempty"`
}

// AddPromoRuleResponse ...
type AddPromoRuleResponse PromoRuleInfoResponse

// UpdatePromoRuleRequest ...
type UpdatePromoRuleRequest struct {
	Description      string  `json:"description,omitempty"`
	Amount           float32 `json:"amount,omitempty"`
	Type             string  `json:"type,omitempty"`
	Target           string  `json:"target,omitempty"`
	Title            string  `json:"title,omitempty"`
	StartsAt         Time    `json:"starts_at,omitempty"`
	EndsAt           Time    `json:"ends_at,omitempty"`
	Enabled          bool    `json:"enabled,omitempty"`
	CreatedAtForeign Time    `json:"created_at_foreign,omitempty"`
	UpdatedAtForeign Time    `json:"updated_at_foreign,omitempty"`
}

// UpdatePromoRuleResponse ...
type UpdatePromoRuleResponse PromoRuleInfoResponse

// DeletePromoRuleResponse ...
type DeletePromoRuleResponse EmptyResponse
