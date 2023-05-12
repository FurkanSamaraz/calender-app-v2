package structures

import (
	"time"

	"github.com/google/uuid"
)

type Employee struct {
	ID         uuid.UUID `db:"id"`
	Name       string    `db:"name"`
	Department string    `db:"department"`
	Email      string    `db:"email"`
	Position   string    `db:"position"`
	HireDate   time.Time `db:"hire_date"`
}

func (e *Employee) TableName() string {
	return "calendar.employee"
}
