package service

import (
	"bookcase/internal/storage"
	"bookcase/models/auth"
	"bookcase/models/author"
	"bookcase/models/book"
	u "bookcase/models/user"
)

type ServiceInterface interface {
	AddAuthor(a author.Author) (int, error)
	GetAuthorList() ([]author.Author, error)
	AddBook(b book.BookAdd) (book.BookAdd, error)
	GetPublishingHouseList() ([]book.PublishingHouse, error)
	GetBookCount() (int, error)
	GetBookList(limit, offset int, sortedBy, sortType string) ([]book.BookUnload, error)
	IsAuthorExists(a author.Author) (bool, error)
	Identify(authData auth.AuthData) (u.User, bool, error)
}

type Service struct {
	ServiceInterface
}

func New(storage storage.StorageInterface) *Service {
	return &Service{
		ServiceInterface: NewService(storage),
	}
}
