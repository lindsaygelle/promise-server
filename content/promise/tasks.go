package promise

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
)

type Tasks []Tasks

func DecodeTasks(readCloser io.ReadCloser) (tasks Tasks, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&tasks)
	if err != nil {
		err = ErrTasks
		return
	}
	return
}

func ScanTasks(scanner interface{ Scan(...interface{}) error }) (tasks Tasks, err error) {
	var b []byte
	err = scanner.Scan(&b)
	if err == sql.ErrNoRows {
		err = ErrTasksNotFound
		return
	}
	tasks, err = DecodeTasks(io.NopCloser(bytes.NewReader(b)))
	return
}
