package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWriteRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, _ := r.BasicAuth()
		fmt.Println(r.Header.Get("Content-Type"))
		require.Equal(t, "/post", r.URL.Path)
		require.Equal(t, "user", user)
		require.Equal(t, "password", pass)
		require.Equal(t, "application/json", r.Header.Get("Content-Type"))
		require.Equal(t, "Spreedly-CLI", r.Header.Get("User-Agent"))

		body, err := ioutil.ReadAll(r.Body)
		require.NoError(t, err)
		require.Equal(t, "{\"AccessKey\":\"user\",\"AccessSecret\":\"password\"}", string(body))
	}))
	defer ts.Close()

	baseURL, _ := url.Parse(ts.URL)
	fmt.Println(ts.URL)
	auth := Auth{
		AccessKey:    "user",
		AccessSecret: "password",
	}
	client := Client{
		BaseURL: baseURL,
		Auth:    auth,
	}
	authByte, _ := json.Marshal(auth)
	resp, err := client.WriteRequest(http.MethodPost, "/post", authByte)
	require.NoError(t, err)
	defer resp.Body.Close()
}

func TestReadRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, _ := r.BasicAuth()
		require.Equal(t, "/get", r.URL.Path)
		require.Equal(t, "user", user)
		require.Equal(t, "password", pass)
		require.Equal(t, "application/json", r.Header.Get("Content-Type"))
		require.Equal(t, "Spreedly-CLI", r.Header.Get("User-Agent"))

		body, err := ioutil.ReadAll(r.Body)
		require.NoError(t, err)
		require.Equal(t, "", string(body))
	}))
	defer ts.Close()

	baseURL, _ := url.Parse(ts.URL)

	auth := Auth{
		AccessKey:    "user",
		AccessSecret: "password",
	}
	client := Client{
		BaseURL: baseURL,
		Auth:    auth,
	}

	resp, err := client.ReadRequest(http.MethodGet, "/get")
	require.NoError(t, err)
	defer resp.Body.Close()
}
