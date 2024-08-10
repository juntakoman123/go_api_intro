package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/juntakoman123/go_api_intro/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	expenseHandler := handler.NewExpenseHandler()

	e.POST("/expenses", expenseHandler.CreateExpense)

	go func() {

		if err := e.Start(":1234"); err != http.ErrServerClosed {
			log.Fatalln("Server closed with error: ", err)
		}

	}()

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)

	defer stop()
	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		panic(fmt.Errorf("failed to graceful shutdown; %w", err))
	}
	log.Println("Server shutdown")

}
