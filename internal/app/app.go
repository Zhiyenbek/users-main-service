package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/Zhiyenbek/users-main-service/config"
	handler "github.com/Zhiyenbek/users-main-service/internal/handler/http"
	"github.com/Zhiyenbek/users-main-service/internal/repository"
	"github.com/Zhiyenbek/users-main-service/internal/repository/connection"
	"github.com/Zhiyenbek/users-main-service/internal/service"
)

func Run() error {
	cfg, err := config.New()
	if err != nil {
		log.Println(err)
		return err
	}
	db, err := connection.NewPostgresDB(cfg.DB)
	if err != nil {
		log.Printf("ERROR: error while creating database: %v", err)
		return err
	}
	defer db.Close()
	repos := repository.New(db, cfg)
	services := service.New(repos, cfg)
	handlers := handler.New(services, cfg)
	srv := http.Server{
		Addr:    ":" + strconv.Itoa(cfg.App.Port),
		Handler: handlers.InitRoutes(),
	}
	errChan := make(chan error, 1)
	go func(errChan chan<- error) {
		log.Printf("server on port: %d have started\n", cfg.App.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println(err)
			errChan <- err
		}
	}(errChan)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-quit:
		log.Println("killing signal was received, shutting down the server")
	case err := <-errChan:
		log.Printf("ERROR: HTTP server error received: %v", err)
	}

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.App.TimeOut)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("WARN: Server forced to shutdown: %v", err)
	}
	return nil

}
