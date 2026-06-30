package StringBuilder

import (
	"strings"
)



func fullNameBuilder(firstName, lastName string) string {
	var fullName strings.Builder
	fullName.WriteString(firstName)
	fullName.WriteString(" ")
	fullName.WriteString(lastName)

	return fullName.String()
}
