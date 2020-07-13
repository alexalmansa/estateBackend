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
	Name      string       `json:"name"`
	Age   int       		`json:"age"`
}
func (i *Renter) Create(conn * pgx.Conn) error {

	now := time.Now()
	row := conn.QueryRow(context.Background(), "INSERT INTO renter (name, age, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id", i.Name, i.Age, now, now)
	err := row.Scan(&i.ID)
	if err != nil{
		fmt.Println(err)
		return fmt.Errorf("There was a problem creating renter")
	}
	return nil
}