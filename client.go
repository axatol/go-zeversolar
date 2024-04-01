package zeversolar

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	// address of inverter web interface
	Address string
	// client to use for requests
	Client Doer
}

func (c *Client) getClient() Doer {
	if c.Client == nil {
		return http.DefaultClient
	}

	return c.Client
}

func (c *Client) GetInverterData(ctx context.Context) (*InverterData, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.Address+"/home.cgi", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %s", err)
	}

	res, err := c.getClient().Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get response: %s", err)
	}

	defer res.Body.Close()
	raw, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %s", err)
	}

	var point InverterData
	if err := point.UnmarshalBinary(raw); err != nil {
		return nil, fmt.Errorf("failed to unmarshal point: %s", err)
	}

	return &point, nil
}
