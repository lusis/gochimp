package api

// https://mailchimp.com/developer/api/marketing/verified-domains/

// ListVerifiedDomainsResponse ...
// https://mailchimp.com/developer/api/marketing/verified-domains/list-sending-domains/
type ListVerifiedDomainsResponse struct {
	Domains []VerifiedDomainInfoResponse `json:"domains"`
	CommonCollectionResponse
}

// VerifiedDomainInfoResponse ...
// https://mailchimp.com/developer/api/marketing/verified-domains/get-domain-info/
type VerifiedDomainInfoResponse struct {
	Domain            string `json:"domain"`
	Verified          bool   `json:"verified"`
	Authenticated     bool   `json:"authenticated"`
	VerificationEmail string `json:"verification_email"`
	VerificationSent  Time   `json:"verification_sent"`
}

// AddVerifiedDomainRequest ...
// https://mailchimp.com/developer/api/marketing/verified-domains/add-domain-to-account/
type AddVerifiedDomainRequest struct {
	VerificationEmail string `json:"verification_email"`
}

// AddVerifiedDomainResponse ...
type AddVerifiedDomainResponse VerifiedDomainInfoResponse

// VerifyDomainRequest ...
// https://mailchimp.com/developer/api/marketing/verified-domains/verify-domain/
type VerifyDomainRequest struct {
	Code string `json:"code"`
}

// VerifyDomainResponse ...
type VerifyDomainResponse VerifiedDomainInfoResponse

// DeleteVerifiedDomainResponse ..
type DeleteVerifiedDomainResponse EmptyResponse
