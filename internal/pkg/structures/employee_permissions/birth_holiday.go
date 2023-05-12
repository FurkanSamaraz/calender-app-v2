package structures

import (
	"time"

	"github.com/google/uuid"
)

type BirthHoliday struct {
	ID          uuid.UUID `db:"id"`
	EmployeeID  uuid.UUID `db:"employee_id"`
	ShiftID     uuid.UUID `db:"shift_id"`
	DateOfBirth time.Time `db:"date_of_birth"`
	StartTime   time.Time `db:"start_time"`
	EndTime     time.Time `db:"end_time"`
}

func (u *BirthHoliday) TableName() string {
	return "calendar.birth_holiday"
}
