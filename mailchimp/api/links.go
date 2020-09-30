package api

// LinkResponse ...
/*
    {
      "rel": "string",
      "href": "string",
      "method": "GET",
      "targetSchema": "string",
      "schema": "string"
	}
*/
type LinkResponse struct {
	Rel          string `json:"rel"`
	Href         string `json:"href"`
	Method       string `json:"method"`
	TargetSchema string `json:"targetSchema"`
	Schema       string `json:"schema"`
}
