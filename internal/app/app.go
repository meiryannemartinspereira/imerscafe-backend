package app

import (
	"net/http"

	"imerscafe-backend/internal/handler"
	"imerscafe-backend/internal/repository"
	"imerscafe-backend/internal/service"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()

	ingredientRepository := repository.NewInMemoryIngredientRepository()
	ingredientService := service.NewIngredientService(ingredientRepository)
	ingredientHandler := handler.NewIngredientHandler(ingredientService)

	mux.HandleFunc("/ingredients", ingredientHandler.Create)

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("imerscafe backend is running\n"))
	})

	return mux
}
