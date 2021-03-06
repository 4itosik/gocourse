package two

type employee struct {
	age int
}

type customer struct {
	age int
}

func maxAgePerson(persons ...interface{}) interface{} {
	var oldest interface{}
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
			oldest = p
		}
	}

	return oldest
}
