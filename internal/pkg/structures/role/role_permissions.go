package structures

import (
	"github.com/google/uuid"
)

type RolePermissions struct {
	RoleID       uuid.UUID `db:"role_id"`
	PermissionID uuid.UUID `db:"permission_id"`
}

func (u *RolePermissions) TableName() string {
	return "calendar.role_permissions"
}
