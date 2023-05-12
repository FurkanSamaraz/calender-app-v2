package structures

import (
	"time"

	"github.com/google/uuid"
)

type ImportantDay struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Date        time.Time `db:"date"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (u *ImportantDay) TableName() string {
	return "calendar.important_day"
}
