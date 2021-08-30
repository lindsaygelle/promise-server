package content

import (
	"time"

	"golang.org/x/text/language"
)

type Content struct {
	Error      error
	Language   language.Tag
	Request    ID
	Status     string
	StatusCode int
	Timestamp  time.Time
}
