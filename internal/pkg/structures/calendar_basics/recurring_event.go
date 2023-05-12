package structures

import (
	"time"

	"github.com/google/uuid"
)

type RecurringEvent struct {
	ID        uuid.UUID `db:"id"`
	Title     string    `db:"title"`
	StartTime time.Time `db:"start_time"`
	EndTime   time.Time `db:"end_time"`
	Frequency string    `db:"frequency"`
	EndDate   time.Time `db:"end_date"`
}

func (u *RecurringEvent) TableName() string {
	return "calendar.recurring_event"
}
