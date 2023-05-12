package structures

import (
	"time"

	"github.com/google/uuid"
)

type AuthProvider struct {
	ID           uuid.UUID `db:"id"`
	ProviderName string    `db:"provider_name"`
	ClientID     string    `db:"client_id"`
	ClientSecret string    `db:"client_secret"`
	Scopes       string    `db:"scopes"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

func (u *AuthProvider) TableName() string {
	return "calendar.auth_provider"
}
