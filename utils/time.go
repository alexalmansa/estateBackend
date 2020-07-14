package utils

import (
	"fmt"
	"time"
)
const (
	layout = "2006-01-02"
)
func ConvertTime(string string)  (time.Time, error){

	t, err := time.Parse(layout, string)

	if err != nil {
		fmt.Println(err)
	}
	return t, err
}
