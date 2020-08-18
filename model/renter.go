package model

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"time"
)

type Renter struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"_"`
	UpdatedAt time.Time `json:"_"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
}

func (i *Renter) Create(conn *pgx.Conn) error {

	now := time.Now()
	row := conn.QueryRow(context.Background(), "INSERT INTO renter (name, age, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id", i.Name, i.Age, now, now)
	err := row.Scan(&i.ID)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("There was a problem creating renter")
	}
	return nil
}

func GetRenters(conn *pgx.Conn, renterId string) ([]Renter, error) {
	var rows pgx.Rows
	var err error
	if renterId != "" {
		rows, err = conn.Query(context.Background(), "SELECT id, name, age, created_at, updated_at FROM renter WHERE id = $1", renterId)

	} else {
		rows, err = conn.Query(context.Background(), "SELECT id, name, age, created_at, updated_at FROM renter ")

	}
	if err != nil {
		fmt.Println(" error getting items %v", err)
		return nil, err
	}
	var renter []Renter
	for rows.Next() {
		i := Renter{}
		err = rows.Scan(&i.ID, &i.Name, &i.Age, &i.CreatedAt, &i.UpdatedAt)
		renter = append(renter, i)
	}
	return renter, nil
}
