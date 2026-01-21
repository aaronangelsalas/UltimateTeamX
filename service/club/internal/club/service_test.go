package club

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/google/uuid"
)

type fakeRepo struct {
	club     Club
	cards    []UserCard
	clubErr  error
	cardsErr error
}

func (f *fakeRepo) GetClubByUserID(_ context.Context, _ uuid.UUID) (Club, error) {
	if f.clubErr != nil {
		return Club{}, f.clubErr
	}
	return f.club, nil
}

func (f *fakeRepo) ListUserCardsByClubID(_ context.Context, _ uuid.UUID) ([]UserCard, error) {
	if f.cardsErr != nil {
		return nil, f.cardsErr
	}
	return f.cards, nil
}

func TestServiceGetMyClubOK(t *testing.T) {
	clubID := uuid.New()
	repo := &fakeRepo{
		club: Club{
			ID:      clubID,
			Credits: 1200,
		},
		cards: []UserCard{
			{ID: uuid.New(), PlayerID: uuid.New(), Locked: false},
			{ID: uuid.New(), PlayerID: uuid.New(), Locked: true},
		},
	}
	service := NewService(repo)

	result, err := service.GetMyClub(context.Background(), uuid.New())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Credits != 1200 {
		t.Fatalf("expected credits 1200, got %d", result.Credits)
	}
	if len(result.Cards) != 2 {
		t.Fatalf("expected 2 cards, got %d", len(result.Cards))
	}
}

func TestServiceGetMyClubNotFound(t *testing.T) {
	repo := &fakeRepo{clubErr: sql.ErrNoRows}
	service := NewService(repo)

	_, err := service.GetMyClub(context.Background(), uuid.New())
	if !errors.Is(err, ErrClubNotFound) {
		t.Fatalf("expected ErrClubNotFound, got %v", err)
	}
}

func TestServiceGetMyClubDBError(t *testing.T) {
	repo := &fakeRepo{clubErr: errors.New("db down")}
	service := NewService(repo)

	_, err := service.GetMyClub(context.Background(), uuid.New())
	if err == nil || errors.Is(err, ErrClubNotFound) {
		t.Fatalf("expected generic error, got %v", err)
	}
}
