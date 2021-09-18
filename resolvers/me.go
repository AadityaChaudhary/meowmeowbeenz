package resolvers

import (
	"encoding/json"
	"net/http"
	"socialcredit/db"
	"strconv"
)

func me(w http.ResponseWriter, req *http.Request) {
	var id []string
	var ok bool
	if id, ok = req.URL.Query()["id"]; !ok || len(id) < 1 {
		http.Error(w, "id missing", http.StatusBadRequest)
		return

	}
	uid, err := strconv.ParseUint(id[0], 0, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var u db.User
	db.DB.First(&u, uid)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&u)
}
