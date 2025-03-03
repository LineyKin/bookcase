package book

import (
	"fmt"
	"strconv"
	"strings"
)

// формат добавления
type BookAdd struct {
	Book
	Name            []LiteraryWork  `json:"name"`
	Author          []int           `json:"author"`
	PublishingHouse PublishingHouse `json:"publishingHouse"`
}

// множественный INSERT в literary_work
func (b BookAdd) GetLWInsertion() string {
	if len(b.Name) == 1 {
		return fmt.Sprintf("('%s', $4)", b.Name[0].Name)
	}

	strList := make([]string, len(b.Name))
	for i := 0; i < len(b.Name); i++ {
		strList[i] = fmt.Sprintf("('%s', $4)", b.Name[i].Name)
	}

	return strings.Join(strList, ",")
}

// массив id авторов для INSERT-запроса в postgres
func (b BookAdd) GetAuthorIdsAsArrayForPG() string {
	if len(b.Author) == 1 {
		return fmt.Sprintf("{%d}", b.Author[0])
	}

	strList := make([]string, len(b.Author))
	for i := 0; i < len(b.Author); i++ {
		strList[i] = strconv.Itoa(b.Author[i])
	}

	return fmt.Sprintf("{%s}", strings.Join(strList, ","))
}

// массив id произведений для INSERT-запроса в postgres
func (b BookAdd) GetLitWorkIdsAsArrayForPG() string {
	if len(b.Name) == 1 {
		return fmt.Sprintf("{%d}", b.Name[0].Id)
	}

	strList := make([]string, len(b.Name))
	for i := 0; i < len(b.Name); i++ {
		strList[i] = strconv.Itoa(b.Name[i].Id)
	}

	return fmt.Sprintf("{%s}", strings.Join(strList, ","))
}

func (b BookAdd) HasAuthors() bool {
	return len(b.Author) != 0
}

// проверка на пустоту списка произведений
func (b BookAdd) IsEmptyNameList() bool {
	for _, lw := range b.Name {
		if !lw.IsEmpty() {
			return false
		}
	}

	return true
}
