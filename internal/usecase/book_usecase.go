package usecase

import (
	"github.com/graphzc/go-clean-arch-example/internal/entity"
	"github.com/graphzc/go-clean-arch-example/internal/repository"
)

type BookUseCase interface {
	GetAllBooks() []entity.Book
}

type bookServiceImpl struct {
	bookRepository repository.BookRepository
}

func NewBookService(bookRepository repository.BookRepository) BookUseCase {
	return &bookServiceImpl{
		bookRepository: bookRepository,
	}
}

func (b *bookServiceImpl) GetAllBooks() []entity.Book {
	return b.bookRepository.GetAllBooks()
}
