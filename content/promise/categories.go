package promise

import "database/sql"

type Categories []Category

func ScanCategories(rows *sql.Rows) (categories Categories, err error) {
	return
}
