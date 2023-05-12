package structures

import (
	"time"

	"github.com/google/uuid"
)

type SpecialHoliday struct {
	ID              uuid.UUID `db:"id"`
	ShiftID         uuid.UUID `db:"shift_id"`
	EmployeeID      uuid.UUID `db:"employee_id"`
	ReasonForExcuse string    `db:"reason_for_excuse"`
	StartTime       time.Time `db:"start_time"`
	EndTime         time.Time `db:"end_time"`
}

func (u *SpecialHoliday) TableName() string {
	return "calendar.special_holiday"
}
