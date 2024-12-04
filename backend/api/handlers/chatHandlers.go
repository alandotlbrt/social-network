package handlers

import (
	"encoding/json"
	"net/http"
	Chat "socialNetwork/api/chat"
)

func loadChat(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	everySaid, err := Chat.RetreiveChat()
	if err {
		json.NewEncoder(w).Encode(everySaid)
	} else {
		json.NewEncoder(w).Encode("error")
	}
}

func loadUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	everySaid, err := Chat.RetreiveUser()
	if err {
		json.NewEncoder(w).Encode(everySaid)
	} else {
		json.NewEncoder(w).Encode("error")
	}
}
