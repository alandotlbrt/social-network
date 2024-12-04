package sessions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"socialNetwork/pkg/db"
	"time"

	"github.com/gofrs/uuid"
)

// this function stock in sessionStore the UUID of the user with his name
// it can return a error
func CreateSession(w http.ResponseWriter, r *http.Request, email string) error {

	_, hasCookie := ValidateSession(w, r)

	if hasCookie {
		return nil
	}

	sessionID, err := uuid.NewV4()
	if err != nil {
		return err
	}

	_, err = db.UpdateInDatabase(nil, "users", "cookie", "email", email, sessionID.String())
	if err != nil {
		return err
	}

	cookie := http.Cookie{
		Name:     "session_token",
		Value:    sessionID.String(),
		Expires:  time.Now().Add(24 * time.Hour),
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		Secure:   false,
	}
	http.SetCookie(w, &cookie)

	return nil
}

// this function destroy an session with the w of the users
func DestroySession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return
	}

	sessionID := cookie.Value

	_, err = db.UpdateInDatabase(nil, "users", "cookie", "cookie", sessionID, "NULL")
	if err != nil {
		fmt.Println(err)
		return
	}
	cookie = &http.Cookie{
		Name:     "session_token",
		Value:    sessionID,
		Expires:  time.Now().Add(-1 * time.Hour),
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		Secure:   false,
	}

	http.SetCookie(w, cookie)
}

// this function is unsing for checking if the user have a cookie or session valid
// it return a boolean ( true if it's good, false if isn't )
func ValidateSession(w http.ResponseWriter, r *http.Request) (string, bool) {

	var reqBody struct {
		SessionToken string `json:"session_token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		return "", false
	}

	if reqBody.SessionToken == "" {
		return "", false
	}

	database, err := db.OpenDB("database")
	if err != nil {
		return "", false
	}
	defer database.Close()

	idUser, err := db.CheckInDatabaseAndGet(database, "users", "cookie", reqBody.SessionToken, "id_user")
	if err != nil {
		return "", false
	}

	if idUser == "" {
		return "", false
	}
	return idUser, true
}

func GetCookie(r *http.Request) string {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return ""
	}
	return cookie.Value
}
