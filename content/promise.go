package content

import "time"

type Promise struct {
	Created time.Time
	Expires time.Time
	ID      ID
	Maker   ID
	Name    string
	Status  ID
	Owner   ID
}
