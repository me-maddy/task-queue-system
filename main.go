package main

import (
	"fmt"
	"net/http"

	"github.com/me-maddy/task-queue-system/src/router"
)

func main() {
	apiRouter := router.SetupRoutes()
	http.Handle("/api/", http.StripPrefix("/api", apiRouter))
	fmt.Println("Service is running on port 8000")
	http.ListenAndServe(":8000", nil)
}
