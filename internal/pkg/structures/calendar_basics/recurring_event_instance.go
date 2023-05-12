package structures

import (
	"time"

	"github.com/google/uuid"
)

type RecurringEventInstance struct {
	ID        uuid.UUID `db:"id"`
	EventID   uuid.UUID `db:"event_id"`
	StartTime time.Time `db:"start_time"`
	EndTime   time.Time `db:"end_time"`
}

func (e *RecurringEventInstance) TableName() string {
	return "calendar.recurring_event_instance"
}
