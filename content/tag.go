package content

import "time"

type Tag struct {
	Content

	Created time.Time
	ID      ID
	Maker   ID
	Promise ID
	Value   string
}
