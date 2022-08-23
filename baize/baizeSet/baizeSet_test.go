package baizeSet

import (
	"fmt"
	"testing"
)

func TestBaizeSet(t *testing.T) {
	s := Set[string]{}
	s.Add("a")
	s.Add("b")
	s.Add("c")
	s.Add("d")
	contains := s.Contains("a")
	fmt.Println(contains)
}
