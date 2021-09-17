package promise

import "database/sql"

type Steps []Step

func ScanSteps(rows *sql.Rows) (steps Steps, err error) {
	return
}
