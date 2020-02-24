package titlecase

import (
	"strings"
)

func TitleCase(str, minor string) string {
	minorWords := make(map[string]struct{})
	for _, s := range strings.Split(minor, " ") {
		minorWords[s] = struct{}{}
	}

	words := strings.Split(strings.ToLower(str), " ")
	var ret strings.Builder
	for i, w := range words {
		_, isMinor := minorWords[w]
		if i == 0 || !isMinor {
			w = strings.Title(w)
		}

		ret.WriteString(w)

		if i != len(words)-1 {
			ret.WriteString(" ")
		}
	}

	return ret.String()
}
