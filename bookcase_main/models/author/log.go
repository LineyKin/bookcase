package author

import (
	"bookcase/models"
	"fmt"
	"time"
)

func (a Author) NewLog() models.UserLog {
	var ul models.UserLog
	ul.Message = fmt.Sprintf("Добавлен автор: %s", a.GetName())
	ul.Timestamp = time.Now()

	return ul
}
