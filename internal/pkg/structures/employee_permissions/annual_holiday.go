package structures

import (
	"time"

	"github.com/google/uuid"
)

type AnnualHoliday struct {
	ID                uuid.UUID `db:"id"`
	ShiftID           uuid.UUID `db:"shift_id"`
	EmployeeID        uuid.UUID `db:"employee_id"`
	AnnualHolidayDays time.Time `db:"annual_holiday_days"`
}

func (u *AnnualHoliday) TableName() string {
	return "calendar.annual_holiday"
}
