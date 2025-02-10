package storage

import (
	"bookcase/internal/db"
	"bookcase/internal/storage/db/postgres"
	"bookcase/internal/storage/db/sqlite"
	"bookcase/models/auth"
	"bookcase/models/author"
	"bookcase/models/book"
	u "bookcase/models/user"
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
	GetUserByAuthLogin(data auth.AuthData) (u.User, error)
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
