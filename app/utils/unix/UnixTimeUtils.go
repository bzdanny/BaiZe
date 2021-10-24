package unix

import (
	"fmt"
	"strconv"
	"time"
)

type Time time.Time

// MarshalJSON implements json.Marshaler.
func (t Time) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("%d", time.Time(t).Unix())
	return []byte(stamp), nil
}

func (t *Time) UnmarshalJSON(b []byte) error {
	parseInt, _ := strconv.ParseInt(string(b), 10, 64)
	*t = Time(time.Unix(parseInt, 0))
	return nil
}
