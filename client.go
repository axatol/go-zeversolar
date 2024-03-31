package zeversolar

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	Address string
}

func (c *Client) GetInverterData(ctx context.Context) (*InverterData, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.Address+"/home.cgi", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %s", err)
	}

	res, err := http.DefaultClient.Do(req)
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
