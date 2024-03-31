package zeversolar_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/axatol/go-zeversolar"
	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	mock := func(w http.ResponseWriter, r *http.Request) { w.Write(rawPoint) }
	ts := httptest.NewServer(http.HandlerFunc(mock))
	client := zeversolar.Client{Address: ts.URL}
	_, err := client.GetInverterData(context.Background())
	require.NoError(t, err)
}
