package tests

// go test -count=1 ./tests

import (
	"bookcase/internal/service"
	"bookcase/internal/storage/db/sqlite"
	"bookcase/lib/env"
	"bookcase/models/author"
	"database/sql"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

type IsAuthorExists struct {
	author       author.Author
	wantIsExists bool
}

// проверка автора на существование в БД
func TestIsAuthorExists(t *testing.T) {
	dbsqlite, err := sql.Open("sqlite", "../"+env.GetDbName())
	if err != nil {
		log.Fatal(err)
	}
	defer dbsqlite.Close()

	Storage := sqlite.New(dbsqlite)
	Service := service.New(Storage)

	testData := []IsAuthorExists{
		{author: author.Author{Name: "Михаил", FatherName: "Афанасьевич", LastName: "Булгаков"}, wantIsExists: true},
		{author: author.Author{Name: "Михаил", LastName: "Булгаков"}, wantIsExists: false},
		{author: author.Author{Name: "Карл", LastName: "Маркс"}, wantIsExists: true},
		{author: author.Author{Name: "Святозар", LastName: "Стёркин"}, wantIsExists: false},
		{author: author.Author{Name: "Вера", FatherName: "Викторовна", LastName: "Камша"}, wantIsExists: true},
		{author: author.Author{Name: "Вера", FatherName: "Викторовна", LastName: "Камша22"}, wantIsExists: false},
	}

	for _, v := range testData {
		IsAuthorExists, err := Service.IsAuthorExists(v.author)
		assert.Equal(t, v.wantIsExists, IsAuthorExists)
		assert.NoError(t, err)
	}
}
