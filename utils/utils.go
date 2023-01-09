package utils

import (
	"time"
)

func GetDt() func() time.Duration {
	lastUpdate := time.Now()
	dt := time.Since(lastUpdate)
	return func() time.Duration {
		dt = time.Since(lastUpdate)
		lastUpdate = time.Now()
		return dt
	}
}
