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
	GetPublishingHouseList() ([]book.PublishingHouse, error)
	GetBookCount(userId int) (int, error)
	GetBookCountTotal() (int, error)
	GetBookList(userId, limit, offset int, sortedBy, sortType string) ([]book.BookUnload, error)
	GetBookListTotal(limit, offset int, sortedBy, sortType string) ([]book.BookUnload, error)
	GetAuthorByName(a author.Author) ([]int, error)
	AuthInterface
	AddBookWithNewPublishingHouse(b *book.BookAdd, userId interface{}) error
	AddBook(b *book.BookAdd, userId interface{}) error
}

type Storage struct {
	StorageInterface
}

func New(db db.AppDB) *Storage {
	return &Storage{
		StorageInterface: postgres.New(db.Connection),
	}
}
