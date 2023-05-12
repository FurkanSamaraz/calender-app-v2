package structures

import (
	"github.com/google/uuid"
)

type AuthPermissions struct {
	ID             uuid.UUID `db:"id"`
	Name           string    `db:"name"`
	AuthProviderID uuid.UUID `db:"auth_provider_id"`
}

func (u *AuthPermissions) TableName() string {
	return "calendar.auth_permissions"
}
