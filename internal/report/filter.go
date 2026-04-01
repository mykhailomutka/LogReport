package report

import "time"

type Filter struct {
	Level string
	Since *time.Time
	Until *time.Time
}
