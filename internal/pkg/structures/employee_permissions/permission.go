package structures

import (
	"time"

	"github.com/google/uuid"
)

type Permission struct {
	ID         uuid.UUID `db:"id"`
	EmployeeID uuid.UUID `db:"employees_id"`
	ShiftID    uuid.UUID `db:"shift_id"`
	Type       string    `db:"type"`
	StartTime  time.Time `db:"start_time"`
	EndTime    time.Time `db:"end_time"`
}

func (u *Permission) TableName() string {
	return "calendar.permission"
}
