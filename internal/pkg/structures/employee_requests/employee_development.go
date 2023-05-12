package structures

import (
	"time"

	"github.com/google/uuid"
)

type EmployeeDevelopment struct {
	ID                   uuid.UUID `db:"id"`
	EmployeeID           uuid.UUID `db:"employee_id"`
	DevelopmentName      string    `db:"development_name"`
	DevelopmentStartDate time.Time `db:"development_start_date"`
	DevelopmentEndDate   time.Time `db:"development_end_date"`
}

func (u *EmployeeDevelopment) TableName() string {
	return "calendar.employee_development"
}
