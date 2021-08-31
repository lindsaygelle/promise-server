package postgres

import "database/sql"

type row struct {
	*sql.Row
}
