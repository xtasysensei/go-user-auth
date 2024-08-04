package user

import (
	"database/sql"
	"fmt"

	"github.com/xtasysensei/go-poll/pkg/models"
)

func GetUserByEmail(email string, db *sql.DB) (*models.User, error) {
	rows, err := db.Query("SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}

	u := new(models.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}

	}

	if u.UserID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}
func IsUsernameTaken(db *sql.DB, username string) (bool, error) {
	var count int
	query := "SELECT COUNT(*) FROM users WHERE username = $1"
	err := db.QueryRow(query, username).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func scanRowIntoUser(rows *sql.Rows) (*models.User, error) {
	user := new(models.User)

	err := rows.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}
func GetUserByID(id int, db *sql.DB) (*models.User, error) {
	rows, err := db.Query("SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	u := new(models.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}

	}

	if u.UserID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}
func GetUserByUsername(db *sql.DB, query_username string) (*models.User, error) {
	rows, err := db.Query("SELECT * FROM users WHERE username = $1", query_username)
	if err != nil {
		return nil, err
	}

	u := new(models.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}

	}

	if u.UserID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func CreateUser(user models.User, db *sql.DB) error {
	query := "INSERT INTO users(username, email, password_hash) VALUES ($1, $2, $3)"
	_, err := db.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}
