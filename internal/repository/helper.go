package repository

import (
	"encoding/base64"
	"time"

	"github.com/cockroachdb/errors"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
)

// DecodeCursor will decode cursor from user for rdb.
func DecodeCursor(encodedTime string) (time.Time, error) {
	byt, err := base64.StdEncoding.DecodeString(encodedTime)
	if err != nil {
		return time.Time{}, errors.Wrap(err, "failed to decode cursor")
	}

	timeString := string(byt)

	t, err := time.Parse(timeFormat, timeString)
	if err != nil {
		return time.Time{}, errors.Wrap(err, "failed to parse time")
	}

	return t, nil
}

// EncodeCursor will encode cursor from rdb to user.
func EncodeCursor(t time.Time) string {
	timeString := t.Format(timeFormat)

	return base64.StdEncoding.EncodeToString([]byte(timeString))
}
