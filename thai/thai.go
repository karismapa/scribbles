package thai

import "unicode"

// IsThai
func IsThai(term string) bool {
	for _, c := range term {
		if unicode.In(c, unicode.Thai) {
			return true
		}
	}

	return false
}
