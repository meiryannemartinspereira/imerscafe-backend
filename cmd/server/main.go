package main

import (
	"fmt"
	"net/http"

	"imerscafe-backend/internal/app"
)

func main() {
	mux := app.New()

	fmt.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}

}
