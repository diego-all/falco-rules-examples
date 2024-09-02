package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

var lock sync.Mutex

func main() {
	http.HandleFunc("/delete-file", deleteFileHandler)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func deleteFileHandler(w http.ResponseWriter, r *http.Request) {
	// Simulamos que obtenemos el nombre del archivo a eliminar de una solicitud
	fileName := r.URL.Query().Get("file")

	// Verificamos si el archivo existe (Time Of Check)
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		http.Error(w, "File does not exist", http.StatusNotFound)
		return
	}

	// Aquí hay una ventana de tiempo en la que un atacante podría eliminar el archivo
	// y crear otro archivo con el mismo nombre antes de que se ejecute el código siguiente.

	// Adquirimos el lock antes de eliminar el archivo para simular una operación prolongada
	lock.Lock()
	defer lock.Unlock()

	// Intentamos leer el archivo antes de eliminarlo
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	// Ahora, eliminamos el archivo (Time Of Use)
	if err := os.Remove(fileName); err != nil {
		http.Error(w, "Error deleting file", http.StatusInternalServerError)
		return
	}

	// Enviamos el contenido del archivo eliminado como respuesta
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}
