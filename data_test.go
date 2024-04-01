package zeversolar_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/axatol/go-zeversolar"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testcases = []struct {
	raw    []byte
	parsed *zeversolar.InverterData
}{
	{
		raw: []byte(strings.Join([]string{
			"1",
			"1",
			"EAB123456789",
			"EFMH6DQ123456789",
			"M11",
			"00B00-111R+22B22-333R",
			"14:05 01/03/2024",
			"OK",
			"1",
			"SX00050123456789",
			"4853",
			"21.80",
			"OK",
			"Error",
		}, "\n")),
		parsed: &zeversolar.InverterData{
			RegistryID:       "EAB123456789",
			RegistryKey:      "EFMH6DQ123456789",
			HardwareVersion:  "M11",
			SoftwareVersion:  "00B00-111R+22B22-333R",
			Timestamp:        time.Date(2024, time.March, 1, 14, 5, 0, 0, time.UTC),
			ZevercloudStatus: "OK",
			SerialNumber:     "SX00050123456789",
			PowerAC:          4853,
			EnergyToday:      21.80,
			Status:           "OK",
		},
	},
	{
		raw: []byte(strings.Join([]string{
			"1",
			"0",
			"000000000000",
			"",
			"",
			"+22B22-333R",
			"00:00 00/00/2000",
			"OK",
			"0",
			"Error",
		}, "\n")),
		parsed: nil,
	},
}

func TestUnmarshalInverterData(t *testing.T) {
	for i, tt := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			actual := new(zeversolar.InverterData)
			err := actual.UnmarshalBinary(tt.raw)
			if tt.parsed == nil {
				require.Error(t, err)
				require.Nil(t, tt.parsed)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.parsed, actual)
			}
		})
	}
}
