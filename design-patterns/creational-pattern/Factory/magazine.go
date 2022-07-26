package factory

import "fmt"

type magazine struct {
	publication
}

// Define a stringer interface that gives a string representation of the type
func (m magazine) String() string {
	return fmt.Sprintf("This is a magazine named %s", m.name)
}

func createMagazine(name string, pages int, publisher string) iPublication {
	return &magazine{
		publication: publication{
			name:      name,
			pages:     pages,
			publisher: publisher,
		},
	}
}
