package adminRouter

import (
	"net/http"

	"github.com/me-maddy/task-queue-system/src/database"
	"github.com/me-maddy/task-queue-system/utils"
)

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	users := database.GetAllUsers()
	utils.SendApiResponse(w, http.StatusOK, true, users)
}

func SetupAdminRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", FindAllUsers)

	return mux
}
