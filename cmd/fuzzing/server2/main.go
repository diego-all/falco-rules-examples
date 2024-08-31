package main

import (
	"fmt"
	"net/http"
	"strings"
)

func vulnerableHandler(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("input")

	// Vulnerabilidad de buffer overflow por la concatenación sin límite del input
	buffer := make([]byte, 1024)
	copy(buffer, input)

	if strings.Contains(string(buffer), "fuzz") {
		fmt.Fprintf(w, "¡Vulnerabilidad encontrada!")
		return
	}

	fmt.Fprintf(w, "Procesado: %s", buffer)
}

func main() {
	http.HandleFunc("/vulnerable", vulnerableHandler)
	fmt.Println("Servidor escuchando en el puerto 8080")
	http.ListenAndServe(":8080", nil)
}
