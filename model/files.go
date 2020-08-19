package model

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
)

func (i *Renter) FileCreate(conn *pgx.Conn, flatId int, filepath string) error {

	row := conn.QueryRow(context.Background(), "INSERT INTO file (flat_id, file_path) VALUES ($1, $2) RETURNING id", flatId, filepath)
	err := row.Scan(&i.ID)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("There was a problem creating file")
	}
	return nil
}
