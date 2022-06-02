package utils

func Tern[T any](cond bool, a T, b T) T {
	if cond {
		return a
	}
	return b
}

// Thanks eatingthenight
// https://stackoverflow.com/a/57213476
func RemoveIndex[T any](s []T, index int) []T {
    ret := make([]T, 0)
    ret = append(ret, s[:index]...)
    return append(ret, s[index+1:]...)
}

// Thanks Eissa N.
// https://stackoverflow.com/a/71289792
func PowInts(x, n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	y := PowInts(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return x * y * y
}
