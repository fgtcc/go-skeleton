package utils

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

type GmtTime struct {
	time.Time
}

func (t GmtTime) MarshalJSON() ([]byte, error) {
	// formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	// return []byte(formatted), nil

	timestamp := strconv.FormatInt(t.Unix(), 10)
	return []byte(timestamp), nil
}

// Value insert timestamp into mysql need this function.
func (t GmtTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *GmtTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = GmtTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
