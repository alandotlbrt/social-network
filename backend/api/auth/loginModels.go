package auth

import (
	"database/sql"
	"fmt"
	"net/http"
	"socialNetwork/pkg/db"

	"golang.org/x/crypto/bcrypt"
)

func LoginMainFunctions(w http.ResponseWriter, r *http.Request, email, password string) (bool, string) {
	var storedPassword string

	db, err := db.OpenDB("database")
	if err != nil {
		fmt.Println(err)
		return false, ""
	}
	defer db.Close()

	err = db.QueryRow("SELECT password FROM users WHERE email = ?", email).Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, "Email not found"
		}
		return false, "An error has occurred"

	}
	if !VerifyPassword(password, storedPassword) {

		return false, "Invalid password"
	}
	return true, "Login successful"
}

// Hashpassword take a string at argument and hashed it.
// it return the password at string, and an error.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// VerifyPassword check if the password noHashed match with the hashed password
// take two arguments, password and hash ( both are string )
// It return a bool, true if they match, false if not.
func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
