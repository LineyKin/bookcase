package service

import (
	"bookcase/internal/storage"
	"bookcase/models/author"
	"bookcase/models/book"
	"errors"
	"fmt"
)

type bookcaseService struct {
	storage storage.StorageInterface
}

func NewService(s storage.StorageInterface) *bookcaseService {
	return &bookcaseService{
		storage: s,
	}
}

func (s *bookcaseService) GetBookListTotal(limit, offset int, sortedBy, sortType string) ([]book.BookUnload, error) {
	return s.storage.GetBookListTotal(limit, offset, sortedBy, sortType)
}

func (s *bookcaseService) GetBookList(userId, limit, offset int, sortedBy, sortType string) ([]book.BookUnload, error) {
	return s.storage.GetBookList(userId, limit, offset, sortedBy, sortType)
}

func (s *bookcaseService) AddBook(b book.BookAdd, userId int) (book.BookAdd, error) {
	// 1. проверяем, чтобы было заполнено хотя бы одно проле с названием
	if b.IsEmptyNameList() {
		return b, errors.New("поле 'Название' не заполнено")
	}

	// 2. проверяем, заполнено ли поле с издательством
	if b.PublishingHouse.IsEmpty() {
		return b, errors.New("поле 'Издательство' не заполнено")
	}

	// 3. в зависимости от того новое ли у нас издательство,
	// либо оно уже было в БД (выбрано из списка в форме)
	// вызывается тот или иной метод.
	if b.PublishingHouse.IsNew() {
		err := s.storage.AddBookWithNewPublishingHouse(&b, userId)
		return b, err
	}

	err := s.storage.AddBook(&b, userId)

	return b, err
}

func (s *bookcaseService) AddAuthor(a author.Author) (int, error) {

	isExists, err := s.IsAuthorExists(a)
	if err != nil {
		return 0, err
	}

	if isExists {
		return 0, fmt.Errorf("автор %s уже добавлен", a.GetName())
	}

	// TODO обработать ошибку
	id, err := s.storage.AddAuthor(a)
	fmt.Println(err)
	return id, nil
}

// проверяем, не заносили ли мы уже этого автора в БД
// исходим из того, что полные тёски среди авторов редкость
func (s *bookcaseService) IsAuthorExists(a author.Author) (bool, error) {
	idList, err := s.storage.GetAuthorByName(a)
	if err != nil {
		return false, err
	}

	len := len(idList)

	if len == 0 {
		return false, nil
	}

	if len == 1 {
		return true, nil
	}

	return true, fmt.Errorf("обнаружены дубликаты автора %s", a.GetName())
}

func (s *bookcaseService) GetAuthorList() ([]author.Author, error) {
	return s.storage.GetAuthorList()
}

func (s *bookcaseService) GetPublishingHouseList() ([]book.PublishingHouse, error) {
	return s.storage.GetPublishingHouseList()
}

func (s *bookcaseService) GetBookCount(userId int) (int, error) {
	return s.storage.GetBookCount(userId)
}

func (s *bookcaseService) GetBookCountTotal() (int, error) {
	return s.storage.GetBookCountTotal()
}
