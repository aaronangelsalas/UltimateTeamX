package service

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"

	"UltimateTeamX/service/identity/internal/repo"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidInput   = errors.New("invalid input")
	ErrEmailExists    = errors.New("email already exists")
	ErrUsernameExists = errors.New("username already exists")
)

// RegisterUser valida input, crea repo.User e chiama UserRepo per salvare
func RegisterUser(idDB *sql.DB, email, username, password string) (string, error) {
	// --- Input Validation ---
	if !isValidEmail(email) {
		return "", fmt.Errorf("%w: invalid email format", ErrInvalidInput)
	}
	if len(username) < 3 {
		return "", fmt.Errorf("%w: username too short", ErrInvalidInput)
	}
	if len(password) < 8 {
		return "", fmt.Errorf("%w: password too short", ErrInvalidInput)
	}

	// --- Hash Password ---
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	// --- 3️⃣ Crea struct User del repo ---
	id := uuid.New().String()
	user := repo.User{
		ID:           id,
		Username:     username,
		Email:        email,
		PasswordHash: string(hashed),
	}

	// --- Save on DB ---
	err = repo.CreateUser(idDB, user)
	if err != nil {
		// Unique violation (e-mail)
		if isUniqueViolation(err, "email") {
			return "", ErrEmailExists
		}
		// Unique violation (username)
		if isUniqueViolation(err, "username") {
			return "", ErrUsernameExists
		}
		return "", fmt.Errorf("failed to create user: %w", err)
	}

	return id, nil
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
	return re.MatchString(email)
}

// Stub per check vincolo unico
func isUniqueViolation(err error, field string) bool {
	// da implementare con pq.Error
	return false
}
