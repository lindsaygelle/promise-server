package content

import "time"

type User struct {
	Content

	Created time.Time
	Email   string
	ID      ID
	Name    string
}
