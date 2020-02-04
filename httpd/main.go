package main

import (
	"CV/http/handler"
	"CV/platform/competence"
	"CV/platform/cursus"
	"CV/platform/experience"
	"CV/platform/projet"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	port := ":3000"

	competence := competence.New()
	cursus := cursus.New()
	experience := experience.New()
	projet := projet.New()

	r := chi.NewRouter()

	r.Get("/competence", handler.CompetenceGet(competence))

	fmt.Println("Le serveur est démarré sur le port : " + port)
	http.ListenAndServe(port, r)
}