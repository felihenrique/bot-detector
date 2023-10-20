package utils

import "time"

func ToIsoString(t time.Time) string {
	return t.Format("2006-01-02T15:04:05")
}
