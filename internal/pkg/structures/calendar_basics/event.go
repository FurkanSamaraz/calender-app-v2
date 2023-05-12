package structures

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID        uuid.UUID `db:"id"`
	Title     string    `db:"title"`
	StartTime time.Time `db:"start_time"`
	EndTime   time.Time `db:"end_time"`
}

func (c *Event) TableName() string {
	return "calendar.event"
}
