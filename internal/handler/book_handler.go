package handler

import (
	"github.com/graphzc/go-clean-arch-example/internal/entity"
	"github.com/graphzc/go-clean-arch-example/internal/usecase"
)

type BookHandler interface {
	GetAllBooks() []entity.Book
}

type bookHandlerImpl struct {
	bookUseCase usecase.BookUseCase
}

func NewBookHandler(bookUseCase usecase.BookUseCase) BookHandler {
	return &bookHandlerImpl{
		bookUseCase: bookUseCase,
	}
}

func (b *bookHandlerImpl) GetAllBooks() []entity.Book {
	return b.bookUseCase.GetAllBooks()
}
