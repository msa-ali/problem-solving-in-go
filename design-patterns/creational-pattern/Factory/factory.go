package factory

import "fmt"

type PublicationType int

const (
	Newspaper = iota
	Magazine
)

// newPublication is a factory function that creates the specified publication
func newPublication(pubType PublicationType, name string, pg int, pub string) (iPublication, error) {
	if pubType == Newspaper {
		return createNewspaper(name, pg, pub), nil
	}
	if pubType == Magazine {
		return createMagazine(name, pg, pub), nil
	}
	return nil, fmt.Errorf("no such publication type")
}
