package handlers

import (
	"fmt"
	"log"
	"net/http"
	"socialNetwork/sessions"
	websocket "socialNetwork/sessions/webSocket"
)

func HandlerMain(port string) {
	mux := http.NewServeMux()
	handlersRedirection(mux)
	handlerWithCORS := EnableCORS(mux)
	fmt.Printf("Server is running on port %s\n", port)
	err := http.ListenAndServe(":"+port, handlerWithCORS)
	if err != nil {
		log.Fatal(err)
	}
}
func handlersRedirection(mux *http.ServeMux) {
	mux.HandleFunc("/login", loginHandler)
	mux.HandleFunc("/register", registerHandler)

	mux.HandleFunc("/api/updatedata", updateHandler)
	mux.HandleFunc("/api/logout", sessions.DestroySession)
	mux.HandleFunc("/api/check", validationHandler)
	mux.HandleFunc("/api/recupdata", userDataHandler)
	mux.HandleFunc("/api/createpost", createPostHandler)
	mux.HandleFunc("/api/loadpost", loadPost)
	mux.HandleFunc("/api/info", otherUserDataHandler)
	mux.HandleFunc("/api/loadChat", loadChat)
	mux.HandleFunc("/api/loadUser", loadUser)

	mux.HandleFunc("/ws", websocket.HandleConnection)

	mux.HandleFunc("/api/follow", followHandlers)
	mux.HandleFunc("/api/unfollow", unFollowHandlers)
	mux.HandleFunc("/api/requestfollow", requestFollowHandler)
	mux.HandleFunc("/api/unrequestfollow", unRequestFollowHandlers)
	mux.HandleFunc("/api/waitingFollowers", waitingFollowersHandlers)
	mux.HandleFunc("/api/declinefollower", declineFollowerHandler)
	mux.HandleFunc("/api/acceptfollower", acceptFollowerHandler)
	mux.HandleFunc("/api/fetchfollowers", fetchfollowersHandler)
	mux.HandleFunc("/api/removefollower", removeFollowerHandler)
}

func EnableCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}
