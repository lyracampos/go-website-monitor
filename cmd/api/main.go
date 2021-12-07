package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	dataimp "website-monitor/internal/infrastructure/data_imp/repositories"

	"github.com/go-openapi/runtime/middleware"

	usecasesimp "website-monitor/internal/application/use_cases_imp"
	"website-monitor/internal/services/api/handlers/health"
	"website-monitor/internal/services/api/handlers/websites"

	"github.com/gorilla/mux"
	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9094", "Bind address for the server")

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

	// website - add
	websiteAddUseCase := usecasesimp.NewWebsiteAddUseCase(data)
	websiteAddHandler := websites.NewWebsiteAddHandler(log, websiteAddUseCase)
	websiteAddRouter := router.Methods(http.MethodPost).Subrouter()
	websiteAddRouter.HandleFunc("/websites", websiteAddHandler.Create)

	// website - edit
	websiteEditUseCase := usecasesimp.NewWebsiteEditUseCase(data)
	websiteEditHandler := websites.NewWebsitesEditHandler(log, websiteEditUseCase)
	websiteEditRouter := router.Methods(http.MethodPut).Subrouter()
	websiteEditRouter.HandleFunc("/websites/{id:[0-9]+}", websiteEditHandler.Edit)

	// website - delete
	websiteDeleteUseCase := usecasesimp.NewWebsiteDeleteUseCase(data)
	websiteDeleteHandler := websites.NewWebsiteDeleteHandler(log, websiteDeleteUseCase)
	websiteDeleteRouter := router.Methods(http.MethodDelete).Subrouter()
	websiteDeleteRouter.HandleFunc("/websites/{id:[0-9]+}", websiteDeleteHandler.Delete)

	// website - enable
	websiteEnableUseCase := usecasesimp.NewWebsiteEnableUseCase(data)
	websiteEnableHandler := websites.NewWebsiteEnableHandler(log, websiteEnableUseCase)
	websiteEnableRouter := router.Methods(http.MethodPut).Subrouter()
	websiteEnableRouter.HandleFunc("/websites/{id:[0-9]+}/enable", websiteEnableHandler.Enable)

	// website - disable
	websiteDisableUseCase := usecasesimp.NewWebsiteDisableUseCase(data)
	websiteDisableHandler := websites.NewWebsiteDisableHandler(log, websiteDisableUseCase)
	websiteDisableRouter := router.Methods(http.MethodPut).Subrouter()
	websiteDisableRouter.HandleFunc("/websites/{id:[0-9]+}/disable", websiteDisableHandler.Disable)

	redocOpts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(redocOpts, nil)
	swaggerRoute := router.Methods(http.MethodGet).Subrouter()
	swaggerRoute.Handle("/docs", sh)
	swaggerRoute.Handle("/swagger.yaml", http.FileServer(http.Dir("./internal/services/api/docs")))

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
