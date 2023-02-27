package utils

func Unpack[T any](to_unpack []T, vars ...*T) {
	for i, current_var := range to_unpack {
		*vars[i] = current_var
	}
}
