package api

// ChimpChatterResponse ...
/*
{
  "chimp_chatter": [
    {
      "title": "1 new subscriber to Your New Campaign!",
      "message": "People are telling their friends about your campaign!",
      "type": "campaigns:forward-to-friend",
      "update_time": "2017-08-04T11:09:01+00:00",
      "url": "http://dev.mailchimp.com/reports/summary?id=1",
      "list_id": "2017-08-04T11:09:01+00:00",
      "campaign_id": "2017-08-04T11:09:01+00:00"
    }
  ],
  "total_items": 0,
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
type ChimpChatterResponse struct {
	ChimpChatter []ChimpChatterItemResponse `json:"chimp_chatter"`
	TotalItems   int32                      `json:"total_items"`
	CommonResponse
}

// ChimpChatterItemResponse ...
type ChimpChatterItemResponse struct {
	Title      string `json:"title"`
	Message    string `json:"message"`
	Type       string `json:"type"`
	UpdateTime Time   `json:"update_time"`
	URL        string `json:"url"`
	ListID     string `json:"list_id"`
	CampaignID string `json:"campaign_id"`
}
