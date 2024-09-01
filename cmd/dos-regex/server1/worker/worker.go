package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

const (
	endpoint          = "http://localhost:8080/update_email"
	emailPayload      = `{"email": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa@malicious-domain-with-very-very-very-long-name.com"}`
	numWorkers        = 100000000 // Número de goroutines concurrentes
	numRequests       = 100000000 // Número total de solicitudes a enviar
	requestsPerWorker = numRequests / numWorkers
)

func main() {
	// WaitGroup para esperar a que todos los workers terminen
	var wg sync.WaitGroup

	// Canal para distribuir las tareas entre los workers
	requests := make(chan int, numRequests)

	// Lanzar los workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, requests, &wg)
	}

	// Enviar tareas (requests) al canal
	for i := 0; i < numRequests; i++ {
		requests <- i
	}

	// Cerrar el canal una vez que todas las tareas han sido enviadas
	close(requests)

	// Esperar a que todos los workers terminen
	wg.Wait()

	fmt.Println("Todas las solicitudes han sido enviadas.")
}

func worker(id int, requests <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	for req := range requests {
		// Enviar la solicitud HTTP
		resp, err := client.Post(endpoint, "application/json", strings.NewReader(emailPayload))
		if err != nil {
			fmt.Printf("Worker %d: Error enviando solicitud %d: %s\n", id, req, err)
			continue
		}

		// Leer la respuesta y cerrarla
		resp.Body.Close()

		fmt.Printf("Worker %d: Solicitud %d enviada, estado: %s\n", id, req, resp.Status)
	}
}
