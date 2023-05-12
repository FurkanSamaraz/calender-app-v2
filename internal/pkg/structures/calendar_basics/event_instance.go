package structures

import (
	"time"

	"github.com/google/uuid"
)

type EventInstance struct {
	ID        uuid.UUID `db:"id"`
	EventID   uuid.UUID `db:"event_id"`
	StartTime time.Time `db:"start_time"`
	EndTime   time.Time `db:"end_time"`
}

func (u *EventInstance) TableName() string {
	return "calendar.event_instance"
}
