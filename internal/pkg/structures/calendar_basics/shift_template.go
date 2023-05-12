package structures

import (
	"time"

	"github.com/google/uuid"
)

type ShiftTemplate struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	StartTime time.Time `db:"start_time"`
	EndTime   time.Time `db:"end_time"`
	DayOfWeek time.Time `db:"day_of_week"`
	Time      time.Time `db:"time"`
}

func (u *ShiftTemplate) TableName() string {
	return "calendar.shift_template"
}
