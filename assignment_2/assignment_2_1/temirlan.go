package assignment_2_1

type temirlan struct {
	Person
}

func newTemirlan() IPerson {
	return &temirlan{
		Person{
			name: "Temirlan",
			age:  18,
		},
	}
}
