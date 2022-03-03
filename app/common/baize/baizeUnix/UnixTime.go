package baizeUnix

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

type BaiZeTime struct {
	time.Time
}

// MarshalJSON implements json.Marshaler.
func (t BaiZeTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	seconds := t.Unix()
	return []byte(strconv.FormatInt(seconds, 10)), nil
}

func (t *BaiZeTime) UnmarshalJSON(b []byte) error {
	parseInt, _ := strconv.ParseInt(string(b), 10, 64)
	t.Time = time.Unix(parseInt, 0)
	return nil
}
func (t BaiZeTime) ToString() string {
	return t.Format("2006-01-02 15:04:05")
}

func (t BaiZeTime) Value() (driver.Value, error) {
	var zeroTime time.Time

	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil

	}

	return t.Time, nil

}

func (t *BaiZeTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)

	if ok {
		*t = BaiZeTime{Time: value}

		return nil

	}

	return fmt.Errorf("can not convert %v to timestamp", v)
}
