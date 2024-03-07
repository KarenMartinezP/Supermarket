package main

import (
	// "encoding/json"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var Information struct{
	FirstName string
	LastName string
}

func main(){

	router:= chi.NewRouter()

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
		
	})

	router.Get("/users/{id}", func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})



	router.Post("/greetings", func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusOK)
		err := json.NewDecoder(r.Body).Decode(&Information)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		println("Hello "+Information.FirstName+" "+ Information.LastName)
	})

	http.ListenAndServe(":8080", router)

}