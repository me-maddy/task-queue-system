package utils

import "time"

func GenerateUniqueId() int64 {
	return time.Now().UnixNano()
}
