package model

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"time"
)

type Building struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"_"`
	UpdatedAt time.Time `json:"_"`
	Name      string       `json:"name"`
	Address   string       `json:"address"`
	Longitude float64       `json:"longitude"`
	Latitude  float64       `json:"latitude"`
}

func (i *Building) Create(conn * pgx.Conn) error {

	now := time.Now()
	row := conn.QueryRow(context.Background(), "INSERT INTO building (name, address, longitude, latitude, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", i.Name, i.Address, i.Longitude, i.Latitude, now, now)
	err := row.Scan(&i.ID)
	if err != nil{
		fmt.Println(err)
		return fmt.Errorf("There was a problem creating building")
	}
	return nil
}
