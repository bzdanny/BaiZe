package unix

import (
	"fmt"
	"strconv"
	"time"
)

type BaiZeTime time.Time

// MarshalJSON implements json.Marshaler.
func (t BaiZeTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("%d", time.Time(t).Unix())
	return []byte(stamp), nil
}

func (t *BaiZeTime) UnmarshalJSON(b []byte) error {
	parseInt, _ := strconv.ParseInt(string(b), 10, 64)
	*t = BaiZeTime(time.Unix(parseInt, 0))
	return nil
}
