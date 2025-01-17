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

	t, err := time.Parse("060102150405", string(bytes))
	if err != nil {
		fmt.Printf("GPSU: Skipping invalid GPS timestamp: %s\n", err)
	} else {
		gpsu.Time = t
	}

	return nil
}
