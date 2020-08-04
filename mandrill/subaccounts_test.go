package mandrill

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSubAccountsList(t *testing.T) {
	resp := `
	[
    {
        "id": "cust-123",
        "name": "ABC Widgets, Inc.",
        "custom_quota": 42,
        "status": "active",
        "reputation": 42,
        "created_at": "2013-01-01 15:30:27",
        "first_sent_at": "2013-01-01 15:30:29",
        "sent_weekly": 42,
        "sent_monthly": 42,
        "sent_total": 42
    }
	]
	`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resp)
	}))
	defer srv.Close()
	client, err := New("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	require.NotNil(t, client)
	res, err := client.SubAccountsList("foo")
	require.NoError(t, err)
	require.Len(t, res, 1)
}
func TestSubAccountsAdd(t *testing.T) {
	resp := `{
		"id": "cust-123",
		"name": "ABC Widgets, Inc.",
		"custom_quota": 42,
		"status": "active",
		"reputation": 42,
		"created_at": "2013-01-01 15:30:27",
		"first_sent_at": "2013-01-01 15:30:29",
		"sent_weekly": 42,
		"sent_monthly": 42,
		"sent_total": 42
	}`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resp)
	}))
	defer srv.Close()
	client, err := New("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	require.NotNil(t, client)
	res, err := client.SubAccountsAdd(SubAccount{})
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, "cust-123", res.ID)
}
func TestSubAccountsInfo(t *testing.T) {
	resp := `{}`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resp)
	}))
	defer srv.Close()
	client, err := New("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	require.NotNil(t, client)
	res, err := client.SubAccountsInfo("foo")
	require.NoError(t, err)
	require.NotNil(t, res)
}
func TestSubAccountsUpdate(t *testing.T) {
	resp := `{}`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resp)
	}))
	defer srv.Close()
	client, err := New("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	require.NotNil(t, client)
	err = client.SubAccountsUpdate(&SubAccount{})
	require.NoError(t, err)
}
func TestSubAccountsDelete(t *testing.T) {
	resp := ``
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resp)
	}))
	defer srv.Close()
	client, err := New("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	require.NotNil(t, client)
	err = client.SubAccountsDelete("foo")
	require.NoError(t, err)
}
func TestSubAccountsPause(t *testing.T) {
	resp := ``
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resp)
	}))
	defer srv.Close()
	client, err := New("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	require.NotNil(t, client)
	err = client.SubAccountsPause("foo")
	require.NoError(t, err)
}
func TestSubAccountsResume(t *testing.T) {
	resp := ``
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resp)
	}))
	defer srv.Close()
	client, err := New("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	require.NotNil(t, client)
	err = client.SubAccountsResume("foo")
	require.NoError(t, err)
}
