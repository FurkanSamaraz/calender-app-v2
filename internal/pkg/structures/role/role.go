package structures

import (
	"github.com/google/uuid"
)

type Role struct {
	ID         uuid.UUID `db:"id"`
	Name       string    `db:"name"`
	EmployeeID string    `db:"employee_id"`
}

func (u *Role) TableName() string {
	return "calendar.role"
}
