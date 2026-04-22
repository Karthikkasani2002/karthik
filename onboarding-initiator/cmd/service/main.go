package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	"onboarding/internal/api"
	"onboarding/internal/config"
	"onboarding/internal/kafka"
	"onboarding/internal/logger"
	"onboarding/internal/metrics"
	"onboarding/internal/postgres"
	"onboarding/internal/redis"
)

func main() {

	cfg := config.Load()

	log := logger.New(cfg)

	log.Info("starting onboarding service")

	db := postgres.New(cfg)

	cache := redis.New(cfg)

	producer := kafka.New(cfg)

	// register prometheus collectors
	metrics.Register()

	handler := api.NewHandler(

		db,
		cache,
		producer,
		log,
	)

	router := mux.NewRouter()

	router.HandleFunc("/onboard", handler.Onboard).Methods("POST")

	router.Handle("/metrics", metrics.Handler())

	router.HandleFunc("/health", handler.Health)

	router.HandleFunc("/ready", handler.Ready)

	srv := &http.Server{

		Addr: cfg.ListenAddr,

		Handler: router,

		ReadTimeout: 10 * time.Second,

		WriteTimeout: 10 * time.Second,

		IdleTimeout: 60 * time.Second,
	}

	go func() {

		log.Info("http server started", "addr", cfg.ListenAddr)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {

			log.Error("server error", "err", err)
		}
	}()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	log.Info("shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	producer.Close()

	db.Close()

	cache.Close()

	srv.Shutdown(ctx)

	log.Info("service stopped")
}
