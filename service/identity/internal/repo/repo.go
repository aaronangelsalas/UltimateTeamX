package repo

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID           string
	Username     string
	Email        string
	PasswordHash string
}

func CreateUser(idDB *sql.DB, user User) error {
	query := `
        INSERT INTO users (id, username, email, password_hash)
        VALUES ($1, $2, $3, $4)
    `
	_, err := idDB.Exec(query, user.ID, user.Username, user.Email, user.PasswordHash)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}
