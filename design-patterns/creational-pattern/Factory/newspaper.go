package factory

import "fmt"

type newspaper struct {
	publication
}

// Define a stringer interface that gives a string representation of the type
func (n newspaper) String() string {
	return fmt.Sprintf("This is a newspaper named %s", n.name)
}

func createNewspaper(name string, pages int, publisher string) iPublication {
	return &newspaper{
		publication: publication{
			name:      name,
			pages:     pages,
			publisher: publisher,
		},
	}
}
