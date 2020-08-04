package mandrill

import (
	"context"
	"time"

	"github.com/lusis/gochimp/mandrill/api"
)

// SubAccount represents a subaccount in mandrill
type SubAccount struct {
	ID             string
	Name           string
	CustomQuota    int32
	Status         string
	Reputation     int32
	CreatedAt      time.Time
	FirstSendAt    time.Time
	SentWeekly     int32
	SentMonthly    int32
	SentTotal      int32
	SentHourly     int32
	Notes          string
	HourlyQuota    int32
	LastThirtyDays Stats
}

// SubAccountsList lists subaccounts
func (c *Client) SubAccountsList(q string) ([]*SubAccount, error) {
	return c.SubAccountsListContext(context.TODO(), q)
}

// SubAccountsListContext lists subaccounts with the provided context
func (c *Client) SubAccountsListContext(ctx context.Context, q string) ([]*SubAccount, error) {
	req := &api.SubAccountsListRequest{}
	resp := &api.SubAccountsListResponse{}
	if err := c.postContext(ctx, "subaccounts/list", req, resp); err != nil {
		return nil, err
	}
	res := []*SubAccount{}
	for _, s := range *resp {
		sa := &SubAccount{
			ID:             s.ID,
			Name:           s.Name,
			CustomQuota:    s.CustomQuota,
			Status:         s.Status,
			Reputation:     s.Reputation,
			CreatedAt:      s.CreatedAt.Time,
			FirstSendAt:    s.FirstSentAt.Time,
			SentWeekly:     s.SentWeekly,
			SentMonthly:    s.SentMonthly,
			SentHourly:     s.SentHourly,
			SentTotal:      s.SentTotal,
			Notes:          s.Notes,
			HourlyQuota:    s.HourlyQuota,
			LastThirtyDays: statsResponseToStats(s.Last30Days),
		}
		res = append(res, sa)
	}

	return res, nil
}

// SubAccountsAdd adds a subaccount
func (c *Client) SubAccountsAdd(a SubAccount) (*SubAccount, error) {
	return c.SubAccountsAddContext(context.TODO(), a)
}

// SubAccountsAddContext adds a subaccount with the provided context
func (c *Client) SubAccountsAddContext(ctx context.Context, a SubAccount) (*SubAccount, error) {
	return nil, nil
}

// SubAccountsInfo returns the info for the subaccount
func (c *Client) SubAccountsInfo(id string) (*SubAccount, error) {
	return c.SubAccountsInfoContext(context.TODO(), id)
}

// SubAccountsInfoContext returns the info for the subaccount with the provided context
func (c *Client) SubAccountsInfoContext(ctx context.Context, id string) (*SubAccount, error) {
	return nil, nil
}

// SubAccountsUpdate updates the subaccount
func (c *Client) SubAccountsUpdate(a *SubAccount) error {
	return c.SubAccountsUpdateContext(context.TODO(), a)
}

// SubAccountsUpdateContext updates the subaccount with the provided context
func (c *Client) SubAccountsUpdateContext(ctx context.Context, a *SubAccount) error {
	return nil
}

// SubAccountsDelete deletes the subaccount
func (c *Client) SubAccountsDelete(id string) error {
	return c.SubAccountsDeleteContext(context.TODO(), id)
}

// SubAccountsDeleteContext deletes the subaccount with the provided context
func (c *Client) SubAccountsDeleteContext(ctx context.Context, id string) error {
	return nil
}

// SubAccountsPause pauses the subaccount
func (c *Client) SubAccountsPause(id string) error {
	return c.SubAccountsPauseContext(context.TODO(), id)
}

// SubAccountsPauseContext pauses the subaccount with the provided context
func (c *Client) SubAccountsPauseContext(ctx context.Context, id string) error {
	return nil
}

// SubAccountsResume resumes the subaccount
func (c *Client) SubAccountsResume(id string) error {
	return c.SubAccountsResumeContext(context.TODO(), id)
}

// SubAccountsResumeContext resumes the subaccount with context
func (c *Client) SubAccountsResumeContext(ctx context.Context, id string) error {
	return nil
}
