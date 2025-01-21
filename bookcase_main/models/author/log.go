package author

import (
	"fmt"
	"time"
)

type AuthorLog struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func (a Author) NewLog() AuthorLog {
	var al AuthorLog
	al.Message = fmt.Sprintf("Добавлен автор: %s, id: %d", a.GetName(), a.Id)
	al.Timestamp = time.Now()

	return al
}
