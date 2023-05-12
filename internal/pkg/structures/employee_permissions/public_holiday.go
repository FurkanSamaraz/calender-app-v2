package structures

import (
	"time"

	"github.com/google/uuid"
)

type PublicHoliday struct {
	ID        uuid.UUID `db:"id"`
	ShiftID   uuid.UUID `db:"shift_id"`
	Name      string    `db:"name"`
	StartTime time.Time `db:"start_time"`
	EndTime   time.Time `db:"end_time"`
}

func (u *PublicHoliday) TableName() string {
	return "calendar.public_holiday"
}
