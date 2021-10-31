package etc

import "strings"

func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
