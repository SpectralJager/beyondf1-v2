package v1

import (
	"beyondf1-v2/golang-api/data"
	"encoding/json"
	"log"
	"net/http"
)

func ApiEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Printf("Hit endpoint: ApiEvents; request from ip: %s", r.RemoteAddr)
	var events = new([]data.Events)
	events, msg := data.GetEvents()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"events": events, "msg": msg})
}
