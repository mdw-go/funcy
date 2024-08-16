package to

import (
	"fmt"
	"strconv"

	"github.com/mdwhatcott/funcy/ranger/is"
)

func String[T any](t T) string {
	switch t := any(t).(type) {
	case string:
		return t
	case fmt.Stringer:
		return t.String()
	case uint8:
		return integerString(t)
	case uint16:
		return integerString(t)
	case uint32:
		return integerString(t)
	case uint64:
		return integerString(t)
	case int8:
		return integerString(t)
	case int16:
		return integerString(t)
	case int32:
		return integerString(t)
	case int64:
		return integerString(t)
	case uintptr:
		return integerString(t)
	case int:
		return integerString(t)
	case uint:
		return integerString(t)
	case float32:
		return strconv.FormatFloat(float64(t), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(t, 'f', -1, 64)
	default:
		return fmt.Sprint(t)
	}
}
func integerString[T is.Integer | uintptr](t T) string {
	return strconv.FormatInt(int64(t), 10)
}
