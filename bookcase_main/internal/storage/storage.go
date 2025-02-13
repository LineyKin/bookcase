package storage

import (
	"bookcase/internal/db"
	"bookcase/internal/storage/db/postgres"
	"bookcase/internal/storage/db/sqlite"
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
	GetBookList(userId, limit, offset int, sortedBy, sortType string) ([]book.BookUnload, error)
	GetTotalBookList(limit, offset int, sortedBy, sortType string) ([]book.BookUnload, error)
	GetAuthorByName(a author.Author) ([]int, error)
	AuthInterface
}

type Storage struct {
	StorageInterface
}

func New(db db.AppDB) *Storage {
	return &Storage{
		StorageInterface: factory(db),
	}
}

func factory(appdb db.AppDB) StorageInterface {
	if appdb.Driver == db.SQLITE_DRIVER {
		return sqlite.New(appdb.Connection)
	}

	return postgres.New(appdb.Connection)
}
