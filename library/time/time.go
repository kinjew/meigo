package time

import (
	"database/sql/driver"
	"fmt"
	gotime "time"

	"meigo/library/log"
)

// MTime 参考 bilibili kratos：https://github.com/bilibili/kratos/blob/master/pkg/time/time.go
type MTime struct {
	gotime.Time
}

func (t MTime) MarshalJSON() ([]byte, error) {
	tune := t.Format(`"2006-01-02 15:04:05"`)
	return []byte(tune), nil
}

// Value insert timestamp into mysql need this function.
func (t MTime) Value() (driver.Value, error) {
	var zeroTime gotime.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *MTime) Scan(v interface{}) error {
	value, ok := v.(gotime.Time)
	if ok {
		*t = MTime{Time: value}
		return nil
	}
	err := fmt.Errorf("can not convert %v to timestamp", v)
	log.Error("MTime error: ", err)
	return err
}
