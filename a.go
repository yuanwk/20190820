package iris

type a struct {
	age int
}

func (a1 *a) set_age(x int) {
	a1.age = x
}

func (a1 *a) get_age() int {
	return a1.age
}
