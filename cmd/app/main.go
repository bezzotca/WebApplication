package main

import (
	"context"
	"log"
	"net/http"

	"webproject/config"
	deliveryhttp "webproject/internal/delivery/http"
	"webproject/internal/delivery/http/handler"
	infrapg "webproject/internal/infrastructure/postgres"
	"webproject/internal/usecase"
	"webproject/pkg/database/postgres"
)

func main() {
	cfg := config.Load()

	ctx := context.Background()

	// Core DB connection — the only place the raw pool is created.
	pool, err := postgres.Connect(ctx, cfg.Postgres)
	if err != nil {
		log.Fatalf("connect postgres: %v", err)
	}
	defer pool.Close()

	// Wiring: infrastructure -> usecase -> delivery.
	userRepo := infrapg.NewUserRepository(pool)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	router := deliveryhttp.NewRouter(userHandler)

	log.Printf("listening on :%s", cfg.HTTPPort)
	if err := http.ListenAndServe(":"+cfg.HTTPPort, router); err != nil {
		log.Fatal(err)
	}
}
