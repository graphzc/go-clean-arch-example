package server

import (
	"net/http"

	"github.com/graphzc/go-clean-arch-example/internal/dto"
	"github.com/graphzc/go-clean-arch-example/internal/handler"
	"github.com/labstack/echo/v4"
)

type bookServer struct {
	bookHandler handler.BookHandler
}

func NewBookServer(bookHander handler.BookHandler) *bookServer {
	return &bookServer{
		bookHandler: bookHander,
	}
}

func (b *bookServer) GetAllBooks(c echo.Context) error {
	bookEntities := b.bookHandler.GetAllBooks()

	bookJSONs := make([]dto.BookJSONDTO, 0, len(bookEntities))
	for _, bookEntity := range bookEntities {
		bookJSONs = append(bookJSONs, dto.BookJSONDTO{
			Name: bookEntity.Name,
		})
	}

	return c.JSON(http.StatusOK, bookJSONs)
}
