package transaction

import "time"

type Transaction struct {
	ID string
	UserID string
	Amount float64
	Type string
	Timestamp time.Time
}