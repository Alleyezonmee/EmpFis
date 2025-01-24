package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Alleyezonmee/EmpFis/internal/database"
	net "github.com/Alleyezonmee/EmpFis/networkresponse"
	"github.com/google/uuid"
)

func (apicfg *ApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name       string `json:"name"`
		Department string `json:"dept"`
		EmpRole    string `json:"role"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		net.RespondWithError(w, 400, "Error parsing json while creating user")
		return
	}

	user, er := apicfg.DB.CreateEmployee(r.Context(), database.CreateEmployeeParams{
		ID:         uuid.New().String(),
		EmpName:    params.Name,
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		Department: params.Department,
		EmpRole:    params.EmpRole,
	})
	if er != nil {
		net.RespondWithError(w, 400, fmt.Sprintf("Couldn't create employee %v", er))
		return
	}

	net.RespondWithJson(w, 201, user)
}
