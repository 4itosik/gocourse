package three

import "io"

func printAll(w io.Writer, values ...interface{}) {
	for _, v := range values {
		if d, ok := v.(string); ok {
			w.Write([]byte(d))
		}
	}
}
