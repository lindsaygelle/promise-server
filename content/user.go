package content

import (
	"time"

	"golang.org/x/text/language"
)

type User struct {
	Content

	Created  time.Time
	Email    string
	ID       ID
	Language language.Tag
	Name     string
}
