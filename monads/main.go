package main

type Functor interface {
	Fmap(func(interface{}) interface{})
}

type Monad interface {
	Sequence(func(interface{}) Monad) Monad
	Return(interface{})
}

type List struct {
	data []interface{}
}

func NewList() *List {
	return &List{
		[]interface{}{},
	}
}

func (l *List) Append(v ...interface{}) {
	l.data = append(l.data, v...)
}

func (l *List) Fmap(f func(interface{}) interface{}) {
	for i, v := range l.data {
		l.data[i] = f(v)
	}
}

func (l *List) Return(v interface{}) {
	l.data = []interface{}{v}
}

func (l *List) Sequence(f func(interface{}) Monad) Monad {
	r := NewList()
	for _, v := range l.data {
		r.Append(f(v).(*List).data...)
	}
	return r
}

// Yup, it's an ugly hack
func (l *List) PrintInts() {
	for _, v := range l.data {
		println(v.(int))
	}
}

func main() {
	l := NewList()
	l.Append(1, 2, 3)
	l.Sequence(func(v interface{}) Monad {
		l := NewList()
		l.Return(v.(int)*5)
		l.Append(v.(int)*7)
		return l
	}).(*List).PrintInts()
}
