package tickets

import (
	"fmt"
	"time"
)

func GenerateTicketCode(prefix string, idOrder int) string {
	date := time.Now().UnixNano()

	codeTicket := fmt.Sprintf("%s-%d-%+v", prefix, idOrder, date)
	return codeTicket
}
