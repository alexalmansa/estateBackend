package model

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"time"
)

type Lease struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"_"`
	UpdatedAt time.Time `json:"_"`
	FlatId    int       `json:"flat_id"`
	RenterId  int       `json:"renter_id"`
	Price     float64   `json:"price"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Deposit   float64   `json:"deposit"`
}

func (i *Lease) Create(conn *pgx.Conn) error {

	now := time.Now()
	row := conn.QueryRow(context.Background(), "INSERT INTO lease (flat_id, renter_id, price, start_date ,end_date, deposit, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", i.FlatId, i.RenterId, i.Price, i.StartDate, i.EndDate, i.Deposit, now, now)
	err := row.Scan(&i.ID)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("There was a problem creating lease")
	}
	return nil
}
