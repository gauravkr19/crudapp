package healthcheck

import (
	"database/sql"
	"net/http"

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
	// Implement your application-specific health checks here
	// Return true if the application is healthy, false otherwise
	return true
}

func isDatabaseHealthy(db *sql.DB) bool {
	// Attempt to connect to the PostgreSQL database
	err := db.Ping()
	if err != nil {
		return false
	}

	return true
}
