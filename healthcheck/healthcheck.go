package healthcheck

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

func HealthzHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Check application-specific health
	if isApplicationHealthy() {
		// Check database connectivity and health
		if isDatabaseHealthy(db) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte("Database is not healthy"))
		}
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("Application is not healthy"))
	}
}

func isApplicationHealthy() bool {
	// Define the URL of your CRUD endpoint
	getEndpoint := "http://localhost:8080/api/records"

	// Create an HTTP client with a timeout to prevent the probe from hanging indefinitely
	httpClient := http.Client{
		Timeout: 5 * time.Second, // Adjust the timeout as needed
	}

	// Attempt to make a GET request to your CRUD endpoint
	response, err := httpClient.Get(getEndpoint)
	if err != nil {
		log.Printf("Error making GET request: %v", err)
		return false // The application is not healthy if the request fails
	}
	defer response.Body.Close()

	// Check the response status code
	if response.StatusCode != http.StatusOK {
		log.Printf("Received non-OK status code: %d", response.StatusCode)
		return false // The application is not healthy if the response status is not OK (200)
	}

	// If the GET request succeeds and returns a status code of 200, consider the application healthy
	// log.Println("Application is healthy")
	return true
}

func isDatabaseHealthy(db *sql.DB) bool {
	// Run a simple query to check database health
	err := db.Ping()
	if err != nil {
		return false
	}

	return true
}
