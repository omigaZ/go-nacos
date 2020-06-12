package utils

import "time"

func TimestampString() string {
	return time.Now().Format("20080103_18:03:08.000")
}
