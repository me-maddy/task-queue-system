package router

import (
	"net/http"

	"github.com/me-maddy/task-queue-system/src/router/adminRouter"
	"github.com/me-maddy/task-queue-system/src/router/user"
	"github.com/me-maddy/task-queue-system/utils"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", HealthHandler)
	mux.Handle("/user/", http.StripPrefix("/user", user.SetupUserRoutes()))
	mux.Handle("/admin/", http.StripPrefix("/admin", adminRouter.SetupAdminRoutes()))

	return mux
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status": "Server is running healthy",
	}
	utils.SendApiResponse(w, http.StatusOK, true, response)
}
