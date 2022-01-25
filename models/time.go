package models

import (
	"encoding/json"
	"time"
)

type Time time.Time

func Now() Time {
	return Time(time.Now())
}

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Format("2006-01-02 15:04:05"))
}

func (t *Time) UnmarshalJSON(data []byte) error {
	tt, err := time.Parse("2006-01-02 15:04:05", string(data))
	if err == nil {
		*t = Time(tt)
		return nil
	}
	return err
}
