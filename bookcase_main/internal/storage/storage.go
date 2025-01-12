package storage

import (
	"bookcase/internal/storage/db/postgres"
	"bookcase/models/author"
	"bookcase/models/book"
	"database/sql"
)

type StorageInterface interface {
	AddAuthor(a author.Author) (int, error)
	GetAuthorList() ([]author.Author, error)
	AddPublishingHouse(phName string) (int, error)
	AddPhysicalBook(b *book.BookAdd) (int, error)
	AddLiteraryWork(lwName string) (int, error)
	LinkBookAndLiteraryWork(lwId, bookId int) error
	LinkAuthorAndLiteraryWork(authorId, bookId int) error
	GetPublishingHouseList() ([]book.PublishingHouse, error)
	GetBookCount() (int, error)
	GetBookList(limit, offset int, sortedBy, sortType string) ([]book.BookUnload, error)
	GetAuthorByName(a author.Author) ([]int, error)
}

type Storage struct {
	StorageInterface
}

func New(db *sql.DB) *Storage {
	return &Storage{
		StorageInterface: postgres.New(db),
	}
}
