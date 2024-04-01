package zeversolar_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/axatol/go-zeversolar"
	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	for i, tt := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			mock := func(w http.ResponseWriter, r *http.Request) { w.Write(tt.raw) }
			ts := httptest.NewServer(http.HandlerFunc(mock))
			client := zeversolar.Client{Address: ts.URL}
			actual, err := client.GetInverterData(context.Background())
			require.NoError(t, err)
			require.Equal(t, tt.parsed, *actual)
		})
	}
}

func TestClientLive(t *testing.T) {
	client := zeversolar.Client{Address: "http://192.168.1.44"}
	actual, err := client.GetInverterData(context.Background())
	require.NoError(t, err)
	require.NotNil(t, actual)
	t.Logf("%+v", *actual)
}
