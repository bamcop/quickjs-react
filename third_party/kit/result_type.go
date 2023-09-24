package kit

type Result[T any] struct {
	value T
	err   error
}

func NewResultV[T any](value T) Result[T] {
	return Result[T]{
		value: value,
		err:   nil,
	}
}

func NewResultE[T any](err error) Result[T] {
	return Result[T]{
		err: err,
	}
}

func (r Result[T]) Ok() bool {
	return r.err == nil
}

func (r Result[T]) Value() (T, error) {
	return r.value, r.err
}

func (r Result[T]) Unwrap() T {
	if r.err != nil {
		panic(r.err)
	}
	return r.value
}
