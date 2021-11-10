package capitalize

import "strings"

func Capitalize(s string) string {
	w := strings.Split(s, " ")
	for i, v := range w {
		cw := strings.ToUpper(string(v[0])) + v[1:]
		w[i] = cw
	}
	return strings.Join(w, " ")
}
