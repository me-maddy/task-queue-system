package user

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/me-maddy/task-queue-system/src/database"
	"github.com/me-maddy/task-queue-system/utils"
)

func registerUser(w http.ResponseWriter, r *http.Request) {
	var input database.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.SendApiResponse(w, http.StatusBadRequest, false, "Invalid Body")
		return
	}

	errorMessages := utils.ValidateData(input)

	if errorMessages != nil {
		utils.SendApiResponse(w, http.StatusBadRequest, false, *errorMessages)
		return
	}

	user := database.User{
		ID:          utils.GenerateUniqueId(),
		Name:        input.Name,
		Email:       input.Email,
		Phonenumber: input.Phonenumber,
		Password:    input.Password,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	database.CreateUser(user)
	response := map[string]any{
		"status": "User registered successfully",
		"user":   user,
	}
	utils.SendApiResponse(w, http.StatusOK, true, response)
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	var input UserLoginInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.SendApiResponse(w, http.StatusBadRequest, false, "Invalid Body")
		return
	}

	errorMessages := utils.ValidateData(input)

	if errorMessages != nil {
		utils.SendApiResponse(w, http.StatusBadRequest, false, *errorMessages)
		return
	}

	user := database.GetUserByEmail(input.Email)

	if user == nil {
		utils.SendApiResponse(w, http.StatusUnauthorized, false, "Invalid email")
		return
	}

	if user.Password != input.Password {
		utils.SendApiResponse(w, http.StatusUnauthorized, false, "Invalid password")
		return
	}

	response := map[string]any{
		"status": "User logged in successfully",
		"user":   user,
	}
	utils.SendApiResponse(w, http.StatusOK, true, response)

}

func SetupUserRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /register", registerUser)
	mux.HandleFunc("POST /login", loginUser)

	return mux
}
