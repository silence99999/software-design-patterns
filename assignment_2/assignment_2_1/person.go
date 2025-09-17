package assignment_2_1

type Person struct {
	name string
	age  int
}

func (p *Person) setName(name string) {
	p.name = name
}

func (p *Person) getName() string {
	return p.name
}

func (p *Person) setAge(age int) {
	p.age = age
}

func (p *Person) getAge() int {
	return p.age
}
