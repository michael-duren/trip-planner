package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handlers) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	u, err := h.Queries.GetUser(r.Context(), 1)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
	}

	resp := make(map[string]string)
	resp["message"] = "Hello World"
	resp["user_email"] = u.Email

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
