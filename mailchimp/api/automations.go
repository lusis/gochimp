package api

// ListAutomationsResponse ...
// https://mailchimp.com/developer/api/marketing/automation/list-automations/
type ListAutomationsResponse struct {
	Automations []AutomationInfoResponse `json:"automations"`
	TotalItems  int32                    `json:"total_items"`
	CommonResponse
}

// AutomationInfoResponse ...
// https://mailchimp.com/developer/api/marketing/automation/get-automation-info/
/*
{
  "id": "string",
  "create_time": "2019-08-24T14:15:22Z",
  "start_time": "2019-08-24T14:15:22Z",
  "status": "save",
  "emails_sent": 0,
  "recipients": {
    "list_id": "string",
    "list_is_active": true,
    "list_name": "string",
    "segment_opts": {
      "saved_segment_id": 0,
      "match": "any",
      "conditions": [
        [
          null
        ]
      ]
    },
    "store_id": "1a2df69xxx"
  },
  "settings": {
    "title": "string",
    "from_name": "string",
    "reply_to": "string",
    "use_conversation": true,
    "to_name": "string",
    "authenticate": true,
    "auto_footer": true,
    "inline_css": true
  },
  "tracking": {
    "opens": true,
    "html_clicks": true,
    "text_clicks": true,
    "goal_tracking": true,
    "ecomm360": true,
    "google_analytics": "string",
    "clicktale": "string",
    "salesforce": {
      "campaign": true,
      "notes": true
    },
    "capsule": {
      "notes": true
    }
  },
  "trigger_settings": {
    "workflow_type": "abandonedBrowse",
    "workflow_title": "string",
    "runtime": {
      "days": [
        "sunday"
      ],
      "hours": {
        "type": "send_asap"
      }
    },
    "workflow_emails_count": 0
  },
  "report_summary": {
    "opens": 0,
    "unique_opens": 0,
    "open_rate": 0,
    "clicks": 0,
    "subscriber_clicks": 0,
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
type AutomationInfoResponse struct {
	ID              string                             `json:"id"`
	CreateTime      Time                               `json:"create_time"`
	StartTime       Time                               `json:"start_time"`
	Status          string                             `json:"status"`
	EmailsSent      int32                              `json:"emails_sent"`
	Recipients      AutomationsRecipientsResponse      `json:"recipients"`
	Settings        AutomationsSettingsResponse        `json:"settings"`
	Tracking        AutomationsTrackingResponse        `json:"tracking"`
	TriggerSettings AutomationsTriggerSettingsResponse `json:"trigger_settings"`
	ReportSummary   AutomationsReportSummaryResponse   `json:"report_summary"`
	CommonResponse
}

// AutomationsSettingsResponse ...
type AutomationsSettingsResponse struct {
	Title           string `json:"title"`
	FromName        string `json:"from_name"`
	ReplyTo         string `json:"reply_to"`
	UseConversation bool   `json:"use_conversation"`
	ToName          string `json:"to_name"`
	Authenticate    bool   `json:"authenticate"`
	AutoFooter      bool   `json:"auto_footer"`
	InlineCSS       bool   `json:"inline_css"`
}

// AutomationsTrackingResponse ...
type AutomationsTrackingResponse struct {
	Opens           bool   `json:"opens"`
	HTMLClicks      bool   `json:"html_clicks"`
	TextClicks      bool   `json:"text_clicks"`
	GoalTracking    bool   `json:"goal_tracking"`
	EComm360        bool   `json:"ecomm360"`
	GoogleAnalytics string `json:"google_analytics"`
	ClickTale       string `json:"clicktale"`
	Salesforce      struct {
		Campaign bool `json:"campaign"`
		Notes    bool `json:"notes"`
	} `json:"salesforce"`
	Capsule struct {
		Notes bool `json:"notes"`
	} `json:"capsule"`
}

// AutomationsRecipientsResponse ...
type AutomationsRecipientsResponse struct {
	ListID       string `json:"list_id"`
	ListIsActive bool   `json:"list_is_active"`
	ListName     string `json:"list_name"`
	SegmentOpts  struct {
		SavedSegmentID int32    `json:"saved_segment_id"`
		Match          string   `json:"match"`
		Conditions     []string `json:"conditions"`
	} `json:"segment_opts"`
	StoreID string `json:"store_id"`
}

// AutomationsReportSummaryResponse ...
type AutomationsReportSummaryResponse struct {
	Opens            int32 `json:"opens"`
	UniqueOpens      int32 `json:"unique_opens"`
	OpenRate         int32 `json:"open_rate"`
	Clicks           int32 `json:"clicks"`
	SubscriberClicks int32 `json:"subscriber_clicks"`
	ClickRate        int32 `json:"click_rate"`
}

// AutomationsTriggerSettingsResponse ...
type AutomationsTriggerSettingsResponse struct {
	WorkflowType  string `json:"workflow_type"`
	WorkflowTitle string `json:"workflow_title"`
	Runtime       struct {
		Days  []string `json:"days"`
		Hours struct {
			Type string `json:"type"`
		} `json:"hours"`
	}
	WorkflowEmailsCount int32 `json:"workflow_emails_count"`
}

// AddAutomationRequest ...
/*
{
  "recipients":
    {
      "list_id":"",
      "store_id":""
    },
  "settings":
    {
      "from_name":"",
      "reply_to":""
    },
  "trigger_settings":
    {
      "workflow_type":""
    }
  }
*/
type AddAutomationRequest struct {
	Recipients struct {
		ListID  string `json:"list_id,omitempty"`
		StoreID string `json:"store_id,omitempty"`
	} `json:"recipients"`
	Settings struct {
		FromName string `json:"from_name,omitempty"`
		ReplyTo  string `json:"reply_to,omitempty"`
	} `json:"settings,omitempty"`
	TriggerSettings struct {
		WorkflowType string `json:"workflow_type"`
	} `json:"trigger_settings"`
}

// UpdateAutomationRequest ...
/*
{
  "settings":{
    "title":"",
    "from_name":"",
    "reply_to":""
  },
  "delay":{
    "amount":0,
    "type":"now",
    "direction":"after",
    "action":"signup"
  }
}
*/
type UpdateAutomationRequest struct {
	Settings struct {
		Title    string `json:"title,omitempty"`
		FromName string `json:"from_name,omitempty"`
		ReplyTo  string `json:"reply_to,omitempty"`
	} `json:"settings,omitempty"`
	Delay struct {
		Amount    int32  `json:"amount,omitempty"`
		Type      string `json:"type,omitempty"`
		Direction string `json:"direction,omitempty"`
		Action    string `json:"action,omitempty"`
	} `json:"delay,omitempty"`
}
