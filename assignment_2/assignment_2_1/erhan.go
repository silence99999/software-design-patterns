package assignment_2_1

type erhan struct {
	Person
}

func newErhan() IPerson {
	return &erhan{
		Person{
			name: "Erhan",
			age:  18,
		},
	}
}
