package api

// EcommerceStoreInfoResponse ...
// https://mailchimp.com/developer/api/marketing/ecommerce-stores/get-store-info/
type EcommerceStoreInfoResponse struct {
	ID            string                         `json:"id"`
	ListID        string                         `json:"list_id"`
	Name          string                         `json:"name"`
	Platform      string                         `json:"platform"`
	Domain        string                         `json:"domain"`
	IsSyncing     bool                           `json:"is_syncing"`
	EmailAddress  string                         `json:"email_address"`
	CurrencyCode  string                         `json:"currency_code"`
	MoneyFormat   string                         `json:"money_format"`
	PrimaryLocale string                         `json:"primary_locale"`
	Timezone      string                         `json:"timezone"`
	Phone         string                         `json:"phone"`
	Address       CommonAddressDetailResponse    `json:"address"`
	ConnectedSite EcommerceConnectedSiteResponse `json:"connected_site"`
	Automations   EcommerceAutomationsResponse   `json:"automations"`
	ListIsActive  bool                           `json:"list_is_active"`
	CreatedAt     Time                           `json:"created_at"`
	UpdatedAt     Time                           `json:"updated_at"`
	CommonResponse
}

// EcommerceConnectedSiteResponse ...
type EcommerceConnectedSiteResponse struct {
	SiteForeignID string `json:"site_foreign_id"`
	SiteScript    struct {
		URL      string `json:"url"`
		Fragment string `json:"fragment"`
	} `json:"site_script"`
}

// EcommerceAutomationsResponse ...
type EcommerceAutomationsResponse struct {
	AbandonedCart   EcommerceAutomationsDetailResponse `json:"abandoned_cart"`
	AbandonedBrowse EcommerceAutomationsDetailResponse `json:"abandoned_browse"`
}

// EcommerceAutomationsDetailResponse ...
type EcommerceAutomationsDetailResponse struct {
	IsSupported bool   `json:"is_supported"`
	ID          string `json:"id"`
	Status      string `json:"status"`
}

// EcommerceAddStoreRequest ...
// https://mailchimp.com/developer/api/marketing/ecommerce-stores/add-store/
type EcommerceAddStoreRequest struct {
	ID            string                     `json:"id"`
	ListID        string                     `json:"list_id"`
	Name          string                     `json:"name"`
	CurrencyCode  string                     `json:"currency_code"`
	Platform      string                     `json:"platform,omitempty"`
	Domain        string                     `json:"domain,omitempty"`
	IsSyncing     bool                       `json:"is_syncing,omitempty"`
	EmailAddress  string                     `json:"email_address,omitempty"`
	MoneyFormat   string                     `json:"money_format,omitempty"`
	PrimaryLocale string                     `json:"primary_locale,omitempty"`
	Timezone      string                     `json:"timezone,omitempty"`
	Phone         string                     `json:"phone,omitempty"`
	Address       CommonAddressDetailRequest `json:"address,omitempty"`
}

// EcommerceAddStoreResponse ...
type EcommerceAddStoreResponse EcommerceStoreInfoResponse

// EcommerceUpdateStoreRequest ...
// https://mailchimp.com/developer/api/marketing/ecommerce-stores/update-store/
type EcommerceUpdateStoreRequest struct {
	Name          string                     `json:"name,omitempty"`
	CurrencyCode  string                     `json:"currency_code,omitempty"`
	Platform      string                     `json:"platform,omitempty"`
	Domain        string                     `json:"domain,omitempty"`
	IsSyncing     bool                       `json:"is_syncing,omitempty"`
	EmailAddress  string                     `json:"email_address,omitempty"`
	MoneyFormat   string                     `json:"money_format,omitempty"`
	PrimaryLocale string                     `json:"primary_locale,omitempty"`
	Timezone      string                     `json:"timezone,omitempty"`
	Phone         string                     `json:"phone,omitempty"`
	Address       CommonAddressDetailRequest `json:"address,omitempty"`
}

// EcommerceUpdateStoreResponse ...
type EcommerceUpdateStoreResponse EcommerceStoreInfoResponse

// EcommerceListStoresResponse ...
// https://mailchimp.com/developer/api/marketing/ecommerce-stores/list-stores/
type EcommerceListStoresResponse struct {
	Stores []EcommerceStoreInfoResponse `json:"stores"`
	CommonCollectionResponse
}

// DeleteStoreResponse ...
type DeleteStoreResponse EmptyResponse
