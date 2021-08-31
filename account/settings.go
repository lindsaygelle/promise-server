package account

import (
	"time"

	"golang.org/x/text/language"
)

type Settings struct {
	Account   uint         `json:"account"`
	Biography string       `json:"biography"`
	Country   uint         `json:"country"`
	Edited    time.Time    `json:"edited"`
	Language  language.Tag `json:"language"`
}
