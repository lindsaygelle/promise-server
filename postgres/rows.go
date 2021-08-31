package postgres

import "database/sql"

type rows struct {
	*sql.Rows
}
