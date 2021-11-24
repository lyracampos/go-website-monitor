package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	usecasesimp "website-monitor/internal/application/use_cases_imp"
	dataimp "website-monitor/internal/infrastructure/data_imp"
	"website-monitor/internal/services/api/handlers/health"
	"website-monitor/internal/services/api/handlers/websites"

	"github.com/gorilla/mux"
	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

// //create a new router
// router := mux.NewRouter()

// //specify endpoints, handler functions and HTTP method
// router.HandleFunc("/health-check", HealthCheck).Methods("GET")
// router.HandleFunc("/persons", Persons).Methods("GET")
// http.Handle("/", router)

// //start and listen to requests
// http.ListenAndServe(":8080", router)

func main() {
	env.Parse()

	router := mux.NewRouter()

	log := log.New(os.Stdout, "websites-monitor-api ", log.LstdFlags)
	repo := dataimp.NewWebsiteData()

	// health
	healthHandler := health.NewHealthHandler(log, repo)
	healthRouter := router.Methods(http.MethodGet).Subrouter()
	healthRouter.HandleFunc("/health-check", healthHandler.Check)

	// website active
	websiteActiveUseCase := usecasesimp.NewWebsiteActiveUseCase(repo)
	websiteActiveHandler := websites.NewWebsiteActiveHandler(log, websiteActiveUseCase)
	websiteActiveRouter := router.Methods(http.MethodPut).Subrouter()
	websiteActiveRouter.HandleFunc("/websites/{id:[0-9]+}/active",websiteActiveHandler.Active)

	http.Handle("/", router)

	s := http.Server{
		Addr:         *bindAddress,
		Handler:      router,
		ErrorLog:     log,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		log.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			log.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
