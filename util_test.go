package zeversolar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestParseTimeDate(t *testing.T) {
	table := []struct {
		input    string
		expected time.Time
	}{
		{"14:05 01/03/2024", time.Date(2024, time.March, 1, 14, 5, 0, 0, time.UTC)},
	}

	for _, tt := range table {
		t.Run(tt.input, func(t *testing.T) {
			actual, err := parseTimeDate(tt.input)
			require.NoError(t, err)
			require.Equal(t, tt.expected, *actual)
		})
	}
}
