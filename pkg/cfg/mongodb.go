package cfg

import "time"

type Mongodb struct {
	Uri      string
	Database string
	Timeout  time.Duration
}
