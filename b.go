package iris

func Main() {
	a1 := a{0}
	age := a1.get_age()
	println(age)
	a1.set_age(24)
	age = a1.get_age()
	println(age)
}
