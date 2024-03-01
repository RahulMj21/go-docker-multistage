package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	fmt.Println("ENV Loaded")
}

func returnJSON(w http.ResponseWriter, status string, message string) {
	resp := Response{
		Status:  status,
		Message: message,
	}

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Error"))
	}

}

func main() {
	loadEnv()
	PORT := os.Getenv("PORT")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		returnJSON(w, "OK", fmt.Sprintf("Hello from : %s", os.Getenv("APP_NAME")))
	})

	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		returnJSON(w, "OK", "Hello World!")
	})

	fmt.Println("server is running on PORT : " + PORT)
	err := http.ListenAndServe(":"+PORT, mux)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
