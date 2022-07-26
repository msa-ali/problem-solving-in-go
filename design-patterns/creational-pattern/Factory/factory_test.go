package factory

import (
	"fmt"
	"testing"
)

func TestFactory(t *testing.T) {
	mag1, _ := newPublication(Magazine, "Tyme", 50, "The Tymes")
	mag2, _ := newPublication(Magazine, "Lyfe", 40, "Lyfe Inc")
	news1, _ := newPublication(Newspaper, "The Herald", 60, "Heralders")
	news2, _ := newPublication(Newspaper, "TOI", 30, "The Time of India")
	pubDetails(mag1)
	pubDetails(mag2)
	pubDetails(news1)
	pubDetails(news2)
}

func pubDetails(pub iPublication) {
	fmt.Printf("--------------------\n")
	fmt.Printf("%s\n", pub)
	fmt.Printf("Type: %T\n", pub)
	fmt.Printf("Name: %s\n", pub.getName())
	fmt.Printf("Pages: %d\n", pub.getPages())
	fmt.Printf("Publisher: %s\n", pub.getPublisher())
	fmt.Printf("--------------------\n")
}
