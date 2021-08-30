package content

import "time"

type Promise struct {
	Created time.Time
	ID      ID
	Maker   ID
	Name    string
	Owner   ID
}
