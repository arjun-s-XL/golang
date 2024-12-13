package main

import (
	"fmt"
	"strings"
)

// getInitials extracts the initials from a full name.
func getInitials(n string) (string, string) {
	// Convert the full name to uppercase
	s := strings.ToUpper(n)
	// Split the name into parts
	names := strings.Split(s, " ")
	var initials []string
	for _, v := range names {
		if len(v) > 0 {
			initials = append(initials, v[:1])
		}
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	} else if len(initials) == 1 {
		return initials[0], ""
	}
	return "", ""
}

func main() {
	fmt.Println(getInitials("Anthony Brigerton"))
}
