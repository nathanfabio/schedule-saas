package repositories

import (
	"log"

	"github.com/nathanfabio/schedule-saas/config"
	"github.com/nathanfabio/schedule-saas/internal/models"
)

// CreateUser creates a new user in the database
func CreateUser(user models.User) error {
	_, err := config.DB.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3, NOW())", user.Name, user.Email, user.Password)
	if err != nil {
		log.Println("Error creating user:", err)
	}
	return err
}

// GetUserByEmail returns a user by email
func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := config.DB.QueryRow("SELECT id, name, email, password FROM users WHERE email = $1", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
