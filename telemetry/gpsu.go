package telemetry

import (
	"errors"
	"fmt"
	"time"
)

// GPS-acquired timestamp
type GPSU struct {
	Time time.Time
}

func (gpsu *GPSU) Parse(bytes []byte) error {
	if 16 != len(bytes) {
		return errors.New("Invalid length GPSU packet")
	}

	str := string(bytes)
	if str == "000000000000.000" {
		fmt.Printf("GPSU: Skipping invalid GPS timestamp: %s\n", str)
	} else {
		t, err := time.Parse("060102150405", string(bytes))
		if err != nil {
			return err
		}
		gpsu.Time = t
	}

	return nil
}
