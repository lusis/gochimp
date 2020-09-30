package api

// ErrorResponse is an error response
/*
{
  "type": "http://developer.mailchimp.com/documentation/mailchimp/guides/error-glossary/",
  "title": "Forbidden",
  "status": 403,
  "detail": "The API key provided is linked to datacenter 'us1'",
  "instance": "XXXXXX"
}
*/
type ErrorResponse struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int32  `json:"status"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}
