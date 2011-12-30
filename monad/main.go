package main

type Functor interface {
	Fmap(func(interface{}) interface{})
}

type List struct {
	data []interface{}
}

func NewList() *List {
	return &List{
		[]interface{}{},
	}
}

func (l *List) Append(v interface{}) {
	l.data = append(l.data, v)
}

func (l *List) Fmap(f func(interface{}) interface{}) {
	for i, v := range l.data {
		l.data[i] = f(v)
	}
}

// Yup, it's an ugly hack
func (l *List) PrintInts() {
	for _, v := range l.data {
		println(v.(int))
	}
}

func main() {
	l := NewList()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Fmap(func(i interface{}) interface{} {
		return i.(int) + 2
	})
	l.PrintInts()
}
