package one

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

type person interface {
	years() int
}

func maxYears(persons ...person) int {
	var years int
	for _, p := range persons {
		if p.years() > years {
			years = p.years()
		}
	}
	return years
}
