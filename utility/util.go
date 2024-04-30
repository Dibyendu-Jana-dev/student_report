package utility

import (
	"log"
	"strings"
)

func RecoverFromPanic(panic string) {
	if r := recover(); r != nil {
		log.Printf("panic: %+v, recovered from %s\n", r, panic)
	}
}

func CheckAscOrDesc(orderBy *string) string {
	if orderBy == nil {
		return "asc"
	}
	order := strings.TrimSpace(*orderBy)
	if strings.EqualFold(order, "desc") {
		return "desc"
	} else if strings.EqualFold(order, "asc") {
		return "asc"
	} else {
		return "asc"
	}
}
