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

func main() {
	env.Parse()

	router := mux.NewRouter()

	log := log.New(os.Stdout, "websites-monitor-api ", log.LstdFlags)
	data := dataimp.NewWebsiteData()

	// health
	healthHandler := health.NewHealthHandler(log, data)
	healthRouter := router.Methods(http.MethodGet).Subrouter()
	healthRouter.HandleFunc("/health-check", healthHandler.Check)

	// webiste - list
	websiteListHandler := websites.NewWebsiteListHandler(log, data)
	websiteRouter := router.Methods(http.MethodGet).Subrouter()
	websiteRouter.HandleFunc("/websites", websiteListHandler.List)

	// website - get
	websiteGetHandler := websites.NewWebsiteGetHandler(log, data)
	websiteGetRouter := router.Methods(http.MethodGet).Subrouter()
	websiteGetRouter.HandleFunc("/websites/{id:[0-9]+}", websiteGetHandler.Get)

	// website - active
	websiteActiveUseCase := usecasesimp.NewWebsiteActiveUseCase(data)
	websiteActiveHandler := websites.NewWebsiteActiveHandler(log, websiteActiveUseCase)
	websiteActiveRouter := router.Methods(http.MethodPut).Subrouter()
	websiteActiveRouter.HandleFunc("/websites/{id:[0-9]+}/active", websiteActiveHandler.Active)

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

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
