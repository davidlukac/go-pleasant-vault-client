package client

import (
	"fmt"
)

func ObfuscatePassword(password string) string {
	var res string

	filler := "*****"
	l := len(password)

	if len(password) > 0 {
		start := string([]rune(password)[0])
		end := string([]rune(password)[l-1])
		res = fmt.Sprintf("%s%s%s", start, filler, end)
	} else {
		res = filler
	}

	return res
}
