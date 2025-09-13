package main

import (
	"github.com/graphzc/go-clean-arch-example/internal/handler"
	"github.com/graphzc/go-clean-arch-example/internal/repository"
	"github.com/graphzc/go-clean-arch-example/internal/server"
	"github.com/graphzc/go-clean-arch-example/internal/usecase"
	"github.com/labstack/echo/v4"
)

func main() {

	bookRepository := repository.NewMemoryBookRepository()
	bookUseCase := usecase.NewBookService(bookRepository)
	bookHandler := handler.NewBookHandler(bookUseCase)
	bookServer := server.NewBookServer(bookHandler)

	e := echo.New()

	e.GET("/books", bookServer.GetAllBooks)

	e.Logger.Fatal(e.Start(":1323"))
}
