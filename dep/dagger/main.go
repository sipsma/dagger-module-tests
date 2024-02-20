package main

type Dep struct{}

func (m *Dep) Fn() string {
	return "hi from dep"
}
