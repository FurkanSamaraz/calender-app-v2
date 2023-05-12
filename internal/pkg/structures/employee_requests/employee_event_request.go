package structures

import (
	"time"

	"github.com/google/uuid"
)

type EmployeeEventRequest struct {
	ID               uuid.UUID `db:"id"`
	EmployeeID       uuid.UUID `db:"employee_id"`
	EventType        string    `db:"event_type"`
	EventDate        time.Time `db:"event_date"`
	EventDescription string    `db:"event_description"`
}

func (e *EmployeeEventRequest) TableName() string {
	return "calendar.employee_event_request"
}
