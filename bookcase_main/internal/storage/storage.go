package storage

import (
	"bookcase/internal/db"
	"bookcase/internal/storage/db/postgres"
	"bookcase/models/author"
	"bookcase/models/book"
)

type StorageInterface interface {
	AddAuthor(a author.Author) (int, error)
	GetAuthorList() ([]author.Author, error)
	AddPublishingHouse(phName string) (int, error)
	AddPhysicalBook(b *book.BookAdd, userId interface{}) (int, error)
	AddLiteraryWork(lwName string) (int, error)
	LinkBookAndLiteraryWork(lwId, bookId int) error
	LinkAuthorAndLiteraryWork(authorId, bookId int) error
	GetPublishingHouseList() ([]book.PublishingHouse, error)
	GetBookCount(userId int) (int, error)
	GetBookCountTotal() (int, error)
	GetBookList(userId, limit, offset int, sortedBy, sortType string) ([]book.BookUnload, error)
	GetBookListTotal(limit, offset int, sortedBy, sortType string) ([]book.BookUnload, error)
	GetAuthorByName(a author.Author) ([]int, error)
	AuthInterface
	AddBookWithNewPublishingHouse(b *book.BookAdd) error
	AddBook(b *book.BookAdd) error
}

type Storage struct {
	StorageInterface
}

func New(db db.AppDB) *Storage {
	return &Storage{
		StorageInterface: postgres.New(db.Connection),
	}
}
