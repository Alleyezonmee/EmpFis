package handlers

import (
	"net/http"

	resp "github.com/Alleyezonmee/EmpFis/networkresponse"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	resp.RespondWithJson(w, 200, struct{}{})
}
