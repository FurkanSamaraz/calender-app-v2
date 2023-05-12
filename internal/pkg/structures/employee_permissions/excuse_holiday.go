package structures

import (
	"time"

	"github.com/google/uuid"
)

type ExcuseHoliday struct {
	ID              uuid.UUID `db:"id"`
	EmployeeID      uuid.UUID `db:"employee_id"`
	ShiftID         uuid.UUID `db:"shift_id"`
	ReasonForExcuse string    `db:"reason_for_excuse"`
	StartTime       time.Time `db:"start_time"`
	EndTime         time.Time `db:"end_time"`
}

func (u *ExcuseHoliday) TableName() string {
	return "calendar.excuse_holiday"
}
