package api

// TODO: timestamp types
// RootResponse ...
/*
{
  "account_id": "string",
  "login_id": "string",
  "account_name": "string",
  "email": "string",
  "first_name": "string",
  "last_name": "string",
  "username": "freddie2000",
  "avatar_url": "string",
  "role": "string",
  "member_since": "2010-01-01T23:59:59+00:00",
  "pricing_plan_type": "monthly",
  "first_payment": "2010-01-01T23:59:59+00:00",
  "account_timezone": "string",
  "account_industry": "string",
  "contact": {
    "company": "string",
    "addr1": "string",
    "addr2": "string",
    "city": "string",
    "state": "string",
    "zip": "string",
    "country": "string"
  },
  "pro_enabled": true,
  "last_login": "2019-08-24T14:15:22Z",
  "total_subscribers": 0,
  "industry_stats": {
    "open_rate": 0,
    "bounce_rate": 0,
    "click_rate": 0
  },
  "_links": [
    {
      "rel": "string",
      "href": "string",
      "method": "GET",
      "targetSchema": "string",
      "schema": "string"
    }
  ]
}
*/
type RootResponse struct {
	AccountID       string `json:"account_id"`
	LoginID         string `json:"login_id"`
	AccountName     string `json:"account_name"`
	Email           string `json:"email"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Username        string `json:"username"`
	AvatarURL       string `json:"avatar_url"`
	Role            string `json:"role"`
	MemberSince     Time   `json:"member_since"`
	PricingPlanType string `json:"pricing_plan_type"`
	FirstPayment    Time   `json:"first_payment"`
	AccountTimezone string `json:"account_timezone"`
	AccountIndustry string `json:"account_industry"`
	Contact         struct {
		Company    string `json:"company"`
		AddressOne string `json:"addr1"`
		AddressTwo string `json:"addr2"`
		City       string `json:"city"`
		State      string `json:"state"`
		Zip        string `json:"zip"`
		Country    string `json:"country"`
	} `json:"contact"`
	ProEnabled       bool  `json:"pro_enabled"`
	LastLogin        Time  `json:"last_login"`
	TotalSubscribers int32 `json:"total_subscribers"`
	IndustryStats    struct {
		OpenRate   int32 `json:"open_rate"`
		BounceRate int32 `json:"bounce_rate"`
		ClickRate  int32 `json:"click_rate"`
	} `json:"industry_stats"`
	Links []LinkResponse `json:"_links"`
}
