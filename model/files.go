package model

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"time"
)

func FileCreate(conn *pgx.Conn, flatId int, filepath string) error {
	now := time.Now()
	row := conn.QueryRow(context.Background(), "INSERT INTO files (flat_id, file_path, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id", flatId, filepath, now, now)
	var id int
	err := row.Scan(&id)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("There was a problem creating file")
	}
	return nil
}

func GetFilesFromFlat(conn *pgx.Conn, flatid string) (ret []string, err error) {
	rows, err := conn.Query(context.Background(), "SELECT file_path FROM files WHERE flat_id = $1", flatid)
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
