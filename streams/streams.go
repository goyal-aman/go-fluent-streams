package streams

import "fmt"

type Supplier[T any] struct {
	Get     func() (*T, bool)
	HasNext func() bool
}

type Stream[T any] struct {
	supplier Supplier[T]
}

func Of[T any](t []T) *Stream[T] {
	var index int = 0

	get := func() (*T, bool) {
		if index >= len(t) {
			return nil, false
		}
		retVal := t[index]
		index++
		return &retVal, index < len(t)
	}
	next := func() bool {
		return index < len(t)
	}
	return &Stream[T]{supplier: Supplier[T]{
		Get:     get,
		HasNext: next,
	}}
}

func OfGenerator[T any](init *T, generator func(T) *T) *Stream[T] {
	curr := init
	get := func() (*T, bool) {
		retVal := generator(*curr)
		curr = retVal
		return retVal, true
	}
	return &Stream[T]{supplier: Supplier[T]{Get: get, HasNext: func() bool { return true }}}

}

func (s *Stream[I]) Map(mapper func(I) *I) *Stream[I] {
	var hasNext bool = true
	f := func() (*I, bool) {
		for itr := s.supplier; itr.HasNext(); {
			val, hasNext := itr.Get()
			return mapper(*val), hasNext
		}
		hasNext = false
		return nil, false
	}
	return &Stream[I]{supplier: Supplier[I]{HasNext: func() bool {
		return hasNext
	},
		Get: f}}
}
func (s *Stream[I]) Filter(predicate func(I) bool) *Stream[I] {
	var hasNext bool = true
	f := func() (*I, bool) {
		for itr := s.supplier; itr.HasNext(); {
			val, hasNext := itr.Get()
			if val != nil && predicate(*val) {
				return val, hasNext
			}
		}
		hasNext = false
		return nil, false
	}
	return &Stream[I]{supplier: Supplier[I]{HasNext: func() bool {
		return hasNext
	},
		Get: f}}
}

func (s *Stream[T]) Collect(collector func(*T)) {
	for itr := s.supplier; itr.HasNext(); {
		get, b := itr.Get()
		collector(get)
		if !b {
			break
		}
	}
}
func (s *Stream[T]) Print() {
	for itr := s.supplier; itr.HasNext(); {
		get, b := itr.Get()
		if get == nil {
			fmt.Println("nil", b)
		} else {
			fmt.Println(*get, b)
		}
		if !b {
			break
		}
	}
}
