package book

import (
	"bookcase/models"
	"fmt"
	"strconv"
	"strings"
)

func (ba BookAdd) NewLog() models.UserLog {
	ul := models.NewUserLog()
	ul.Message = fmt.Sprintf(
		"Добавлена книга '%s', ID автора(ов): %s; издательство %s, год издания %s",
		ba.GetName(),
		ba.GetAuthors(),
		ba.PublishingHouse.Name,
		ba.PublishingYear,
	)

	return ul
}

func (b BookAdd) GetAuthors() string {
	author := make([]string, len(b.Author))
	for i := 0; i < len(b.Author); i++ {
		author[i] = strconv.Itoa(b.Author[i])
	}

	return strings.Join(author, ",")
}

func (b BookAdd) GetName() string {
	name := make([]string, len(b.Name))
	for i := 0; i < len(b.Name); i++ {
		name[i] = b.Name[i].Name
	}

	return strings.Join(name, ". ")
}
