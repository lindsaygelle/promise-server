package account

import "time"

type Account struct {
	Created time.Time `json:"created"`
	ID      uint      `json:"id"`
	Name    string    `json:"name"`
}
