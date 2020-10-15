package api

// VerifiedDomainsResponse ...
/*
{
  "domains": [
    {
      "domain": "string",
      "verified": true,
      "authenticated": true,
      "verification_email": "string",
      "verification_sent": "2019-08-24T14:15:22Z"
    }
  ],
  "total_items": 0
}
*/
type VerifiedDomainsResponse struct {
	Domains    []VerifiedDomainInfoResponse `json:"domains"`
	TotalItems int32                        `json:"total_items"`
}

// VerifiedDomainInfoResponse ...
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
