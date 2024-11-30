package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/brnocorreia/api-golang-template/internal/api"
	"github.com/brnocorreia/api-golang-template/internal/config"
	"github.com/brnocorreia/api-golang-template/internal/store/pgstore"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	ctx := context.Background()

	queries := pgstore.InitDB(ctx)
	defer pgstore.CloseDB()
	pgstore.MigrateDB(ctx)

	handler := api.NewHandler(queries)

	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", config.ApiPort), handler); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				panic(err)
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}