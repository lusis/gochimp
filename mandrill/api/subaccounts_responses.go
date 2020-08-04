package api

type SubAccountsInfoResponse struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Notes       string       `json:"notes"`
	CustomQuota int32        `json:"custom_quota"`
	Status      string       `json:"status"`
	Reputation  int32        `json:"reputation"`
	CreatedAt   Time         `json:"created_at"`
	FirstSentAt Time         `json:"first_sent_at"`
	SentWeekly  int32        `json:"sent_weekly"`
	SentMonthly int32        `json:"sent_monthly"`
	SentTotal   int32        `json:"sent_total"`
	SentHourly  int32        `json:"sent_hourly"`
	HourlyQuota int32        `json:"hourly_quota"`
	Last30Days  StatResponse `json:"last_30_days"`
}

type SubAccountsListResponse []SubAccountsInfoResponse

/*
type SubAccountsListResponse []struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CustomQuota int    `json:"custom_quota"`
	Status      string `json:"status"`
	Reputation  int    `json:"reputation"`
	CreatedAt   Time   `json:"created_at"`
	FirstSentAt Time   `json:"first_sent_at"`
	SentWeekly  int    `json:"sent_weekly"`
	SentMonthly int    `json:"sent_monthly"`
	SentTotal   int    `json:"sent_total"`
}
*/
type SubAccountsAddResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CustomQuota int    `json:"custom_quota"`
	Status      string `json:"status"`
	Reputation  int    `json:"reputation"`
	CreatedAt   Time   `json:"created_at"`
	FirstSentAt Time   `json:"first_sent_at"`
	SentWeekly  int    `json:"sent_weekly"`
	SentMonthly int    `json:"sent_monthly"`
	SentTotal   int    `json:"sent_total"`
}

type SubAccountsUpdateResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CustomQuota int    `json:"custom_quota"`
	Status      string `json:"status"`
	Reputation  int    `json:"reputation"`
	CreatedAt   Time   `json:"created_at"`
	FirstSentAt Time   `json:"first_sent_at"`
	SentWeekly  int    `json:"sent_weekly"`
	SentMonthly int    `json:"sent_monthly"`
	SentTotal   int    `json:"sent_total"`
}

type SubAccountsDeleteResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CustomQuota int    `json:"custom_quota"`
	Status      string `json:"status"`
	Reputation  int    `json:"reputation"`
	CreatedAt   Time   `json:"created_at"`
	FirstSentAt Time   `json:"first_sent_at"`
	SentWeekly  int    `json:"sent_weekly"`
	SentMonthly int    `json:"sent_monthly"`
	SentTotal   int    `json:"sent_total"`
}

type SubAccountsPauseResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CustomQuota int    `json:"custom_quota"`
	Status      string `json:"status"`
	Reputation  int    `json:"reputation"`
	CreatedAt   Time   `json:"created_at"`
	FirstSentAt Time   `json:"first_sent_at"`
	SentWeekly  int    `json:"sent_weekly"`
	SentMonthly int    `json:"sent_monthly"`
	SentTotal   int    `json:"sent_total"`
}

type SubAccountsResumeResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CustomQuota int    `json:"custom_quota"`
	Status      string `json:"status"`
	Reputation  int    `json:"reputation"`
	CreatedAt   Time   `json:"created_at"`
	FirstSentAt Time   `json:"first_sent_at"`
	SentWeekly  int    `json:"sent_weekly"`
	SentMonthly int    `json:"sent_monthly"`
	SentTotal   int    `json:"sent_total"`
}
