package mandrill

import (
	"context"

	"github.com/lusis/gochimp/mandrill/api"
)

// URL represents the data about a given url in Mandrill
type URL struct {
	URL          string
	Sent         int32
	Clicks       int32
	UniqueClicks int32
}

// ListURLS lists all URLs
func (c *Client) ListURLS() ([]URL, error) {
	return c.ListURLSContext(context.TODO())
}

// ListURLSContext lists all URLs
func (c *Client) ListURLSContext(ctx context.Context) ([]URL, error) {
	req := &api.URLsListRequest{}
	resp := &api.URLsListResponse{}
	err := c.postContext(ctx, "urls/list", req, resp)
	if err != nil {
		return nil, err
	}
	urls := make([]URL, len(*resp))
	for _, u := range *resp {
		urls = append(urls, URL{
			URL:          u.URL,
			Sent:         u.Sent,
			Clicks:       u.Clicks,
			UniqueClicks: u.UniqueClicks,
		})
	}
	return urls, nil
}

// SearchURLS searches for urls using the provided query
func (c *Client) SearchURLS(q string) ([]URL, error) {
	return c.SearchURLSContext(context.TODO(), q)
}

// SearchURLSContext searches for urls using the provided query
func (c *Client) SearchURLSContext(ctx context.Context, q string) ([]URL, error) {
	req := &api.URLsSearchRequest{
		Q: q,
	}
	resp := &api.URLsSearchResponse{}
	err := c.postContext(ctx, "urls/list", req, resp)
	if err != nil {
		return nil, err
	}
	urls := make([]URL, len(*resp))
	for _, u := range *resp {
		urls = append(urls, URL{
			URL:          u.URL,
			Sent:         u.Sent,
			Clicks:       u.Clicks,
			UniqueClicks: u.UniqueClicks,
		})
	}
	return urls, nil
}

// GetURLTimeSeries returns time-series faceted stats about the provided url
func (c *Client) GetURLTimeSeries(url string) ([]TimeSeries, error) {
	return c.GetURLTimeSeriesContext(context.TODO(), url)
}

// GetURLTimeSeriesContext returns time-series faceted stats about the provided url
func (c *Client) GetURLTimeSeriesContext(ctx context.Context, url string) ([]TimeSeries, error) {
	req := &api.URLsTimeSeriesRequest{
		URL: url,
	}
	resp := &api.URLsTimeSeriesResponse{}
	err := c.postContext(ctx, "urls/time-series", req, resp)
	if err != nil {
		return nil, err
	}
	tsData := make([]TimeSeries, len(*resp))
	for _, t := range *resp {
		tsData = append(tsData, TimeSeries{
			Time:         t.Time.Time,
			Sent:         t.Sent,
			Clicks:       t.Clicks,
			UniqueClicks: t.UniqueClicks,
		})
	}
	return tsData, nil
}

// AddTrackingDomain adds a new tracking domain
func (c *Client) AddTrackingDomain(domain string) (*TrackingDomain, error) {
	return c.AddTrackingDomainContext(context.TODO(), domain)
}

// AddTrackingDomainContext adds a new tracking domain
func (c *Client) AddTrackingDomainContext(ctx context.Context, domain string) (*TrackingDomain, error) {
	req := &api.URLsAddTrackingDomainRequest{
		Domain: domain,
	}
	resp := &api.URLsAddTrackingDomainResponse{}
	err := c.postContext(ctx, "urls/add-tracking-domain", req, resp)
	if err != nil {
		return nil, err
	}
	return &TrackingDomain{
		Domain:       resp.Domain,
		CreatedAt:    resp.CreatedAt.Time,
		LastTestedAt: resp.LastTestedAt.Time,
		CName: CName{
			Valid:      resp.CName.Valid,
			ValidAfter: resp.CName.ValidAfter.Time,
			Error:      resp.CName.Error,
		},
		ValidTracking: resp.ValidTracking,
	}, nil
}

// CheckTrackingDomain checks the status of the provided tracking domain
func (c *Client) CheckTrackingDomain(domain string) (*TrackingDomain, error) {
	return c.CheckTrackingDomainContext(context.TODO(), domain)
}

// CheckTrackingDomainContext checks the status of the provided tracking domain
func (c *Client) CheckTrackingDomainContext(ctx context.Context, domain string) (*TrackingDomain, error) {
	req := &api.URLsAddTrackingDomainRequest{
		Domain: domain,
	}
	resp := &api.URLsAddTrackingDomainResponse{}
	err := c.postContext(ctx, "urls/add-tracking-domain", req, resp)
	if err != nil {
		return nil, err
	}
	return &TrackingDomain{
		Domain:       resp.Domain,
		CreatedAt:    resp.CreatedAt.Time,
		LastTestedAt: resp.LastTestedAt.Time,
		CName: CName{
			Valid:      resp.CName.Valid,
			ValidAfter: resp.CName.ValidAfter.Time,
			Error:      resp.CName.Error,
		},
		ValidTracking: resp.ValidTracking,
	}, nil
}
