package api

import (
	"strings"
	"time"
)

const (
	DateFormate     = "2006-01-02"
	DateTimeFormate = "2006-01-02 15:04:05"
)

type JsonDate time.Time
type JsonDateTime time.Time

func (tt *JsonDate) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == "null" {
		return nil
	}
	text := strings.ReplaceAll(string(data), "\"", "")
	d, err := time.ParseInLocation(DateFormate, text, time.Local)
	*tt = (JsonDate)(d)
	return err
}

func (tt JsonDate) MarshalJSON() ([]byte, error) {
	t := time.Time(tt)
	return []byte("\"" + t.Format(DateFormate) + "\""), nil
}

func (date *JsonDateTime) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == "null" {
		return nil
	}
	text := strings.ReplaceAll(string(data), "\"", "")
	d, err := time.ParseInLocation(DateTimeFormate, text, time.Local)
	*date = JsonDateTime(d)
	return err
}

func (date JsonDateTime) MarshalJSON() ([]byte, error) {
	t := time.Time(date)
	return []byte("\"" + t.Format(DateTimeFormate) + "\""), nil
}
