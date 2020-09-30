package api

// ListAuthorizedAppsResponse ...
// https://mailchimp.com/developer/api/marketing/authorized-apps/list-authorized-apps/
/*
{
  "apps": [
    {
      "id": 0,
      "name": "string",
      "description": "string",
      "users": [
        "string"
      ],
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
type ListAuthorizedAppsResponse struct {
	CommonResponse
	TotalItems int32                       `json:"total_items"`
	Apps       []AuthorizedAppInfoResponse `json:"apps"`
}

// AuthorizedAppInfoResponse ...
/*
{
  "id": 0,
  "name": "string",
  "description": "string",
  "users": [
    "string"
  ],
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
type AuthorizedAppInfoResponse struct {
	ID          int32    `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Users       []string `json:"users"`
	CommonResponse
}

// LinkApplicationResponse ...
// https://mailchimp.com/developer/api/marketing/authorized-apps/link-application/
/*
{
  "access_token": "string",
  "viewer_token": "string",
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
type LinkApplicationResponse struct {
	AccessToken string `json:"access_token"`
	ViewerToken string `json:"viewer_token"`
	CommonResponse
}

// LinkApplicationRequest ...
type LinkApplicationRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}
