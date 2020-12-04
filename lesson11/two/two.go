package two

type employee struct {
	age int
}

type customer struct {
	age int
}

func (e employee) years() int {
	return e.age
}

func (c customer) years() int {
	return c.age
}

func maxAgePerson(persons ...interface{}) interface{} {
	var oldperson interface{}
	var years int

	for _, p := range persons {
		var age int

		if u, ok := p.(employee); ok {
			age = u.age
		}
		if u, ok := p.(customer); ok {
			age = u.age
		}

		if age > years {
			years = age
			oldperson = p
		}
	}

	return oldperson
}
