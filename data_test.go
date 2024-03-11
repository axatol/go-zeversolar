package zeversolar_test

import (
	"testing"
	"time"

	"github.com/axatol/go-zeversolar"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var rawPoint = []byte(`1
1
EAB123456789
EFMH6DQ123456789
M11
00B00-111R+22B22-333R
14:05 01/03/2024
OK
1
SX00050123456789
4853
21.80
OK
Error

`)

func TestUnmarshalInverterData(t *testing.T) {
	expected := zeversolar.InverterData{
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
	}

	var actual zeversolar.InverterData
	err := actual.UnmarshalBinary(rawPoint)
	require.NoError(t, err)
	assert.Equal(t, expected, actual)
}
