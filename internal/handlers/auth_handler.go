package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/nathanfabio/schedule-saas/internal/models"
	"github.com/nathanfabio/schedule-saas/repositories"
	"golang.org/x/crypto/bcrypt"
)

// Credencials is a login request struct
type Credencials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterUser registers a new user
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error to read request body", http.StatusBadRequest)
		return
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error to hash password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hash)

	// Save on database
	err = repositories.CreateUser(user)
	if err != nil {
		http.Error(w, "Error to register user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}


// LoginUser logs in a user
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var credencials Credencials

	err := json.NewDecoder(r.Body).Decode(&credencials)
	if err != nil {
		http.Error(w, "Error to read request body", http.StatusBadRequest)
		return
	}

	// Get user from database
	user, err := repositories.GetUserByEmail(credencials.Email)
	if err != nil {
		http.Error(w, "User or password incorrect", http.StatusUnauthorized)
		return
	} 

	// Compare password with hash
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credencials.Password))
	if err != nil {
		http.Error(w, "User or password incorrect", http.StatusUnauthorized)
		return
	}

	// Create JWT token
	expirationTime := time.Now().Add(24 * time.Hour) //valid for 24 hours
	claims := &jwt.StandardClaims{
		Subject:   user.Email,
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		http.Error(w, "Error to create token", http.StatusInternalServerError)
		return
	}

	// Return token
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})

}