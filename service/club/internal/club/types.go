package club

import (
	"context"

	"github.com/google/uuid"
)

// MyClubReader espone l'accesso al club dell'utente corrente.
type MyClubReader interface {
	GetMyClub(ctx context.Context, userID uuid.UUID) (*MyClub, error)
}

// MyClub rappresenta il club dell'utente con i dati necessari al dominio.
type MyClub struct {
	Credits int64
	Cards   []UserCard
}

// UserCard rappresenta una carta posseduta dal club.
type UserCard struct {
	ID       uuid.UUID
	PlayerID uuid.UUID
	Locked   bool
}
