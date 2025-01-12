package book

type Book struct {
	Id             int    `json:"id,omitempty"`
	PublishingYear string `json:"publishingYear"`
}

// формат добавления
type BookAdd struct {
	Book
	Name            []LiteraryWork  `json:"name"`
	Author          []int           `json:"author"`
	PublishingHouse PublishingHouse `json:"publishingHouse"`
}

// формат чтения из списка
type BookUnload struct {
	Book
	Name            string `json:"name"`
	Author          string `json:"author"`
	PublishingHouse string `json:"publishingHouse"`
}

type ListParam struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
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
