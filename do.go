package do

func Run[T any](fn func() T) T {
	return fn()
}
