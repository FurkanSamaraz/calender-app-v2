package structures

import (
	"time"

	"github.com/google/uuid"
)

type Shift struct {
	ID         uuid.UUID `db:"id"`
	TemplateID uuid.UUID `db:"template_id"`
	EmployeeID uuid.UUID `db:"employee_id"`
	Name       string    `db:"name"`
	StartTime  time.Time `db:"start_time"`
	EndTime    time.Time `db:"end_time"`
}

func (u *Shift) TableName() string {
	return "calendar.shift"
}
