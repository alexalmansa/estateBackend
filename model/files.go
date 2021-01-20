package model

import (
	"database/sql"
	"fmt"
	"time"
)

func FileCreate(conn *sql.DB, flatId int, filepath string) error {
	now := time.Now()
	row := conn.QueryRow("INSERT INTO files (flat_id, file_path, created_at, updated_at) VALUES (?,?,?,?)", flatId, filepath, now, now)
	var id int
	err := row.Scan(&id)
	if err != sql.ErrNoRows {
		fmt.Println(err)
		return fmt.Errorf("There was a problem creating file")
	}
	return nil
}

func GetFilesFromFlat(conn *sql.DB, flatid string) (ret []string, err error) {
	rows, err := conn.Query("SELECT file_path FROM files WHERE flat_id = ?", flatid)
	if err != nil {
		fmt.Println(" error getting items %v", err)
		return nil, err
	}
	var filename = ""
	for rows.Next() {

		err = rows.Scan(&filename)
		ret = append(ret, filename)
	}
	return ret, nil
}
