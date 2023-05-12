package structures

import (
	"time"

	"github.com/google/uuid"
)

type Expense struct {
	ID          uuid.UUID `db:"id"`
	Type        string    `db:"type"`
	Description string    `db:"description"`
	Amount      float64   `db:"amount"`
	Date        time.Time `db:"date"`
}

func (u *Expense) TableName() string {
	return "calendar.expense"
}
