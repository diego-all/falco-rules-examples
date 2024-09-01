package main

import (
	"encoding/json"
	"net/http"
	"regexp"
)

type RequestData struct {
	Email string `json:"email"`
}

func updateEmail(w http.ResponseWriter, r *http.Request) {
	var requestData RequestData

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	email := requestData.Email

	// Expresión regular vulnerable a un ataque de DoS
	match, _ := regexp.MatchString(
		`^([0-9a-zA-Z]([-.\w]*[0-9a-zA-Z])*@{1}([0-9a-zA-Z][-\w]*[0-9a-zA-Z]\.)+[a-zA-Z]{2,9})$`,
		email,
	)

	if match {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Correo actualizado correctamente"}`))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "Correo inválido"}`))
	}
}

func main() {
	http.HandleFunc("/update_email", updateEmail)

	http.ListenAndServe(":8080", nil)
}
