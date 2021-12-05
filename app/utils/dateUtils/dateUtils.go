package dateUtils

import "time"

func DatePath() string {
	now := time.Now()
	return now.Format("/2006/01/02")
}
