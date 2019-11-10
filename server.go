package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type apiResponse struct {
	Status   int           `json:"status"`
	Msg      string        `json:"msg"`
	Response greetResponse `json:"response"`
}

type greetResponse struct {
	Greet string    `json:"greeting"`
	Date  time.Time `json:"date"`
}

func main() {
	http.HandleFunc("/", greeter)
	http.HandleFunc("/v1/greet", jsonGreeter)
	http.HandleFunc("/v1/post", post)
	server := &http.Server{Addr: ":8001"}
	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}

func greeter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func jsonGreeter(w http.ResponseWriter, r *http.Request) {
	apiRes := apiResponse{http.StatusOK, "", greetResponse{"Hello json!", time.Now()}}

	res, err := json.Marshal(apiRes)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(res)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}

}

func post(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		param := r.FormValue("key")
		apiRes := apiResponse{http.StatusOK, "", greetResponse{"Hello! Your param is " + param, time.Now()}}

		res, err := json.Marshal(apiRes)
		if err != nil {
			http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(res)
		if err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
		}

	default:
		http.Error(w, "Not support method", http.StatusBadRequest)
	}
}
