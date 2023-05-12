package structures

import (
	"time"

	"github.com/google/uuid"
)

type Notifications struct {
	ID         uuid.UUID `db:"id"`
	EmployeeID uuid.UUID `db:"employees_id"`
	Message    string    `db:"message"`
	CreatedAt  time.Time `db:"created_at"`
}
