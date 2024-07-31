package to

import (
	"fmt"
	"strconv"
)

func String[T any](t T) string {
	switch t := any(t).(type) {
	case fmt.Stringer:
		return t.String()
	case int:
		return strconv.Itoa(t)
	// TODO: other integer types
	default:
		return fmt.Sprint(t)
	}
}
