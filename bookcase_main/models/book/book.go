package book

type Book struct {
	Id             int    `json:"id,omitempty"`
	PublishingYear string `json:"publishingYear"`
}

// формат чтения из списка
type BookUnload struct {
	Book
	Name            string `json:"name"`
	Author          string `json:"author"`
	PublishingHouse string `json:"publishingHouse"`
	User            string `json:"user"`
}

type ListParam struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
