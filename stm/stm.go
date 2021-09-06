package stm

import "strings"

func StringToStringToSliceOfString(s string) map[string][]string {
	m := make(map[string][]string)

	for _, scScplit := range strings.Split(s, ";") {
		cl := strings.Split(scScplit, ":")

		if len(cl) > 1 {
			m[cl[0]] = strings.Split(cl[1], ",")
		}
	}

	return m
}
