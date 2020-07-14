package model

import (
	"context"
	"estateBackend/utils"
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
	StartDate string    `json:"start_date"`
	EndDate   string    `json:"end_date"`
	Deposit   float64   `json:"deposit"`
}

func (i *Lease) Create(conn *pgx.Conn) error {

	now := time.Now()
	startDate, errStart := utils.ConvertTime(i.StartDate)
	endDate, errEnd := utils.ConvertTime(i.EndDate)
	if errStart != nil || errEnd != nil {
		fmt.Println("Error with date format")
	} else {
		row := conn.QueryRow(context.Background(), "INSERT INTO lease (flat_id, renter_id, price, start_date ,end_date, deposit, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", i.FlatId, i.RenterId, i.Price, startDate, endDate, i.Deposit, now, now)
		err2 := row.Scan(&i.ID)
		if err2 != nil {
			fmt.Println(err2)
			return fmt.Errorf("There was a problem creating lease")
		}
	}
	return nil

}
