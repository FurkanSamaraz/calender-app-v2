package structures

import (
	"time"

	"github.com/google/uuid"
)

type EmployeeTraining struct {
	ID                uuid.UUID `db:"id"`
	EmployeeID        uuid.UUID `db:"employee_id"`
	TrainingName      string    `db:"training_name"`
	TrainingStartDate time.Time `db:"training_start_date"`
	TrainingEndDate   time.Time `db:"training_end_date"`
	CertificateName   string    `db:"certificate_name"`
}

func (u *EmployeeTraining) TableName() string {
	return "calendar.employee_training"
}
