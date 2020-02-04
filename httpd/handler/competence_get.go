package handler

import (
	"CV/platform/competence"
	"encoding/json"
	"net/http"
)

func CompetenceGet(comp competence.Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items := comp.GetAll()
		json.NewEncoder(w).Encode(items)
	}
}