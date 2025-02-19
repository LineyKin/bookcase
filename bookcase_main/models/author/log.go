package author

import (
	"bookcase/models"
	"fmt"
)

func (a Author) NewLog() models.UserLog {
	ul := models.NewUserLog()
	ul.Message = fmt.Sprintf("Добавлен автор: %s", a.GetName())

	return ul
}
