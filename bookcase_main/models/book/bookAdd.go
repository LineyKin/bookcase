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
