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

	recipeRepository := repository.NewInMemoryRecipeRepository()
	recipeService := service.NewRecipeService(recipeRepository, ingredientRepository)
	recipeHandler := handler.NewRecipeHandler(recipeService)

	mux.HandleFunc("POST /ingredients", ingredientHandler.Create)
	mux.HandleFunc("GET /ingredients", ingredientHandler.GetAll)
	mux.HandleFunc("GET /ingredients/{id}", ingredientHandler.GetByID)

	mux.HandleFunc("POST /recipes", recipeHandler.Create)
	mux.HandleFunc("GET /recipes", recipeHandler.GetAll)
	mux.HandleFunc("GET /recipes/{id}", recipeHandler.GetByID)

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("imerscafe backend is running\n"))
	})

	return mux
}
