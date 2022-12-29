package funcy

func Complement[T any](predicate func(t T) bool) func(t T) bool {
	return func(t T) bool { return !predicate(t) }
}
func Remove[T any](predicate func(t T) bool, values []T) []T {
	return Filter(Complement(predicate), values)
}
func Filter[T any](predicate func(t T) bool, values []T) (result []T) {
	for _, value := range values {
		if predicate(value) {
			result = append(result, value)
		}
	}
	return result
}
func Map[I, O any](transform func(i I) O, values []I) (result []O) {
	for _, value := range values {
		result = append(result, transform(value))
	}
	return result
}
func MapCat[I, O any](transform func(i I) []O, values []I) (result []O) {
	for _, value := range values {
		result = append(result, transform(value)...)
	}
	return result
}
func Reduce[T any](calc func(a, b T) T, start T, values []T) (result T) {
	result = start
	for _, next := range values {
		result = calc(result, next)
	}
	return result
}
func Range[N Number](start, stop N) (result []N) {
	for start < stop {
		result = append(result, start)
		start++
	}
	return result
}
func TakeWhile[T any](predicate func(T) bool, values []T) (result []T) {
	for _, value := range values {
		if !predicate(value) {
			break
		}
		result = append(result, value)
	}
	return result
}
func DropWhile[T any](predicate func(T) bool, values []T) (result []T) {
	for _, value := range values {
		if predicate(value) {
			continue
		}
		result = append(result, value)
	}
	return result
}

func Add[T Number](a, b T) T        { return a + b }
func Multiply[T Number](a, b T) T   { return a * b }
func Subtract[T Number](a, b T) T   { return a - b }
func Square[T Number](t T) T        { return t * t }
func IsEven[T Integer](t T) bool    { return t%2 == 0 }
func IsOdd[T Integer](t T) bool     { return t%2 == 1 }
func IsPositive[T Number](t T) bool { return t > 0 }
func IsNegative[T Number](t T) bool { return t < 0 }
func IsZero[T Number](t T) bool     { return t == 0 }

type (
	Number interface {
		Integer | uintptr | float64 | float32
	}
	Integer interface {
		int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
	}
)
