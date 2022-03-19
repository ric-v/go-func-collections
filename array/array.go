package array

// arrTypes is a list of all the types that can be used in arrays.
type arrTypes interface {
	~bool | ~uint | ~uint16 | ~uint32 | ~uint64 | ~int8 | ~int | ~int16 | ~int64 | ~float32 | ~float64 | ~complex64 | ~complex128 | ~uintptr | ~string | ~rune | ~byte
}

// In returns the index of the first element in the array that equals the searchKey.
func In[T arrTypes](arr []T, key T) (exists bool, i int, val T) {
	return in(arr, key)
}

// private func for In
func in[T arrTypes](arr []T, key T) (exists bool, i int, val T) {
	for i, val = range arr {
		if val == key {
			return true, i, val
		}
	}
	return false, i, key
}

// Replace replaces the first element in the array that equals the searchKey with the newValue.
func Replace[T arrTypes](arr []T, key T, replacer T) []T {
	_, i, _ := in(arr, key)
	arr[i] = replacer
	return arr
}

// Map applies the function f to each element of the array and returns the result.
func Map[T arrTypes](arr []T, f func(T) T) []T {
	for i, v := range arr {
		arr[i] = f(v)
	}
	return arr
}

// Reduce applies the function f to each element of the array and returns the result.
func Reduce[T arrTypes](a []T, f func(T, T) T) T {
	for i := 1; i < len(a); i++ {
		a[i] = f(a[i-1], a[i])
	}
	return a[len(a)-1]
}

// Filter returns a new array containing only the elements that satisfy the predicate f.
func Filter[T arrTypes](a []T, f func(T) bool) []T {
	var b []T
	for _, v := range a {
		if f(v) {
			b = append(b, v)
		}
	}
	return b
}

// RemoveElem removes the first element in the array that equals the searchKey.
func RemoveElem[T arrTypes](a []T, searchKey T) []T {
	_, i, _ := in(a, searchKey)
	// check if i is out of bounds
	if i+1 > len(a) {
		return a[:i]
	}
	return append(a[:i], a[i+1:]...)
}

// RemoveAt removes the element at the index i.
func RemoveAt[T arrTypes](a []T, i int) []T {
	if i+1 > len(a) {
		return a[:i]
	}
	return append(a[:i], a[i+1:]...)
}

// InsertAt inserts the newValue at the index i.
func InsertAt[T arrTypes](a []T, i int, v T) []T {
	// check if i is out of bounds
	if i > len(a) {
		return append(a, v)
	}
	return append(a[:i], append([]T{v}, a[i:]...)...)
}

// AppendAt appends the newValue at the index i.
func AppendAt[T arrTypes](a []T, i int, v []T) []T {
	// check if i is out of bounds
	if i > len(a) {
		return append(a, v...)
	}
	return append(a[:i], append(v, a[i:]...)...)
}
