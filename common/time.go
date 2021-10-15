package common

import "time"

func TimeNow() string {
	return time.Now().Format("2006-01-02_15:04:05.000")
}
