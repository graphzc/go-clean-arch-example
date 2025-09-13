package repository

import "github.com/graphzc/go-clean-arch-example/internal/entity"

type BookModel struct {
	Name string
}

type BookRepository interface {
	GetAllBooks() []entity.Book
}

type memoryBookRepositoryImpl struct {
	Books []BookModel
}

func NewMemoryBookRepository() BookRepository {
	return &memoryBookRepositoryImpl{
		[]BookModel{
			{
				Name: "Clean Architecture",
			},
			{
				Name: "Clean Code",
			},
		},
	}
}

func (m *memoryBookRepositoryImpl) GetAllBooks() []entity.Book {
	bookEntities := make([]entity.Book, 0, len(m.Books))

	for _, bookModel := range m.Books {
		bookEntities = append(bookEntities, entity.Book{
			Name: bookModel.Name,
		})
	}

	return bookEntities
}
