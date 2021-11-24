package main

import (
	"log"
	"os"
	usecasesimp "website-monitor/internal/application/use_cases_imp"
	dataimp "website-monitor/internal/infrastructure/data_imp"
	"website-monitor/internal/services/api/handlers"
)

func main() {
	//env.Parse()
	// configuration := config.New()
	// database := config.NewMongoDatabase(configuration)

	// // Setup Repository
	// productRepository := repository.NewProductRepository(database)

	// // Setup Service
	// productService := service.NewProductService(&productRepository)

	// // Setup Controller
	// productController := controller.NewProductController(&productService)

	l := log.New(os.Stdout, "products-api ", log.LstdFlags)
	repo := dataimp.NewWebsiteData()
	usecase := usecasesimp.NewWebsiteActiveUseCase(repo)
	handler := handlers.NewWebsiteActiveHandler(l, usecase)

	// activeUseCase := handlers.NewWebsiteActiveHandler(l, usecases.WebsiteActiveUseCase)

	// // create the handlers
	// ph := handlers.NewWebsiteActiveHandler(l)

	// // create a new serve mux and register the handlers
	// sm := mux.NewRouter()

	// getRouter := sm.Methods(http.MethodGet).Subrouter()
	// getRouter.HandleFunc("/", ph.GetProducts)

	// putRouter := sm.Methods(http.MethodPut).Subrouter()
	// putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	// putRouter.Use(ph.MiddlewareValidateProduct)

	// postRouter := sm.Methods(http.MethodPost).Subrouter()
	// postRouter.HandleFunc("/", ph.AddProduct)
	// postRouter.Use(ph.MiddlewareValidateProduct)

}
