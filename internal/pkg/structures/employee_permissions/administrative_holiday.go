package structures

import (
	"time"

	"github.com/google/uuid"
)

type AdministrativeHoliday struct {
	ID        uuid.UUID `db:"id"`
	ShiftID   uuid.UUID `db:"shift_id"`
	Name      string    `db:"name"`
	StartTime time.Time `db:"start_time"`
	EndTime   time.Time `db:"end_time"`
}

func (u *AdministrativeHoliday) TableName() string {
	return "calendar.administrative_holiday"
}
