package gcloud_functions

import (
	"fmt"
	"net/http"
)

func TriggerHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprint(w, "Cloud Function より Hello World (GET)")
	case http.MethodPost:
		fmt.Fprint(w, "Cloud Function より Hello World(POST)")
	default:
		http.Error(w, "Method Not Allowd", http.StatusMethodNotAllowed)
	}
}
