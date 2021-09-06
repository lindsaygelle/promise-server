package database

import (
	"database/sql"
	"encoding/json"
)

type NullTime struct {
	sql.NullTime
}

func (n NullTime) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.Time)
	}
	return []byte(null), nil
}
