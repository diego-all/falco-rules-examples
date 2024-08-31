package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var buffer [10]byte

	// Lee toda la solicitud en el búfer sin verificar el tamaño
	_, err := r.Body.Read(buffer[:])
	if err != nil {
		http.Error(w, "Error leyendo solicitud", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Datos recibidos: %s", buffer)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
