package structures

import (
	"time"

	"github.com/google/uuid"
)

type Send_Emails struct {
	ID         uuid.UUID `db:"id"`
	EmployeeID uuid.UUID `db:"employees_id"`
	Subject    string    `db:"subject"`
	Body       string    `db:"body"`
	CreatedAt  time.Time `db:"created_at"`
}
