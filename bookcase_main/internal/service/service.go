package service

import (
	"bookcase/internal/storage"
	"bookcase/models/author"
	"bookcase/models/book"
)

type ServiceInterface interface {
	AddAuthor(a author.Author) (int, error)
	GetAuthorList() ([]author.Author, error)
	AddBook(b book.BookAdd, userId int) (book.BookAdd, error)
	GetPublishingHouseList() ([]book.PublishingHouse, error)
	GetBookCount(userId int) (int, error)
	GetBookList(userId, limit, offset int, sortedBy, sortType string) ([]book.BookUnload, error)
	IsAuthorExists(a author.Author) (bool, error)
	AuthInterface
}

type Service struct {
	ServiceInterface
}

func New(storage storage.StorageInterface) *Service {
	return &Service{
		ServiceInterface: NewService(storage),
	}
}
