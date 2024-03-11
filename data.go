package zeversolar

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// InverterData is the response from the inverter
//
// data looks something like
//
//	index value                 // description
//	0     1                     // ?
//	1     1                     // ?
//	2     EAB123456789          // registry id
//	3     EFMH6DQ123456789      // registry key
//	4     M11                   // hardware version
//	5     16B00-111R+22B33-444R // software version
//	6     14:05 01/03/2024      // timestamp
//	7     OK                    // zevercloud status
//	8     1                     // ?
//	9     SX00050123456789      // serial number
//	10    4853                  // power ac
//	11    21.80                 // energy today
//	12    OK                    // status
//	13    Error                 // ?
type InverterData struct {
	RegistryID       string
	RegistryKey      string
	HardwareVersion  string
	SoftwareVersion  string
	Timestamp        time.Time
	ZevercloudStatus string
	SerialNumber     string
	PowerAC          int64
	EnergyToday      float64
	Status           string
}

func (d *InverterData) UnmarshalBinary(data []byte) error {
	fields := strings.Split(strings.TrimSpace(string(data)), "\n")
	for i, field := range fields {
		fields[i] = strings.TrimSpace(field)
	}

	if len(fields) != 14 {
		return fmt.Errorf("invalid number of fields: %d", len(fields))
	}

	timestamp, err := parseTimeDate(fields[6])
	if err != nil {
		return fmt.Errorf("failed to parse timestamp %s: %s", fields[6], err)
	}

	powerAC, err := strconv.ParseInt(fields[10], 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse power ac %s: %s", fields[10], err)
	}

	energyToday, err := strconv.ParseFloat(fields[11], 64)
	if err != nil {
		return fmt.Errorf("failed to parse energy today %s: %s", fields[11], err)
	}

	*d = InverterData{
		RegistryID:       fields[2],
		RegistryKey:      fields[3],
		HardwareVersion:  fields[4],
		SoftwareVersion:  fields[5],
		Timestamp:        *timestamp,
		ZevercloudStatus: fields[7],
		SerialNumber:     fields[9],
		PowerAC:          powerAC,
		EnergyToday:      energyToday,
		Status:           fields[12],
	}

	return nil
}
