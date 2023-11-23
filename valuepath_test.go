package filter

import (
	"fmt"
	"testing"
)

func ExampleParseValuePath() {
	fmt.Println(ParseValuePath([]byte("emails[type eq \"work\"]")))
	fmt.Println(ParseValuePath([]byte("emails[not (type eq \"work\")]")))
	fmt.Println(ParseValuePath([]byte("emails[type eq \"work\" and value co \"@example.com\"]")))
	// Output:
	// emails[type eq "work"] <nil>
	// emails[not (type eq "work")] <nil>
	// emails[type eq "work" and value co "@example.com"] <nil>
}

func TestValuePath(t *testing.T) {
	e1 := "emails[type eq \"work\"]"
	e2 := "emails[not (type eq \"work\")]"
	e3 := "emails[type eq \"work\" and value co \"example.com\" or value co \"test.com\"]"
	vp1, _ := ParseValuePath([]byte(e1))
	vp2, _ := ParseValuePath([]byte(e2))
	vp3, _ := ParseValuePath([]byte(e3))
	fmt.Printf("e1: %s", vp1.String())
	fmt.Printf("e2: %s", vp2.String())
	fmt.Printf("e3: %s", vp3.String())
}
