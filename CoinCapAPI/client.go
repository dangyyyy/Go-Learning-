package Coincap

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Client struct {
	client *http.Client
}

func NewClient(time time.Duration) (*Client, error) {
	if time == 0 {
		return nil, errors.New("time must be greater than zero")
	}
	return &Client{
		client: &http.Client{
			Timeout: time,
			Transport: &loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
	}, nil
}
func (c Client) GetAssets() ([]Asset, error) {
	resp, err := c.client.Get("https://api.coincap.io/v2/assets")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err

	}
	var r assetsResponse
	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}
	return r.Assets, nil
}
func (c Client) GetAsset(name string) (Asset, error) {
	url := fmt.Sprintf("https://api.coincap.io/v2/assets/%s", name)
	resp, err := c.client.Get(url)
	if err != nil {
		return Asset{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Asset{}, err

	}
	var r assetResponse
	if err := json.Unmarshal(body, &r); err != nil {
		return Asset{}, err
	}
	return r.Asset, nil
}
