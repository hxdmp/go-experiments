// This program serves as a technical proof that the String method implemented with a value receiver on type X
// effectively supports formatting via fmt.Stringer interface both for value and pointer types. The implementation
// demonstrates compatibility across direct value references and pointers to the type, ensuring that fmt.Printf
// correctly invokes the String method in various contexts. 
package main

import (
	"fmt"
)

type X struct {
	y int
}

func (x X) String() string {
	return fmt.Sprintf("%d", x.y)
}

func main() {
	var (
		vx X  = X{42}
		px *X = &X{69}
	)

	fmt.Printf("%s\n", vx)
	fmt.Printf("%s\n", &vx)
	fmt.Printf("%s\n", px)
	fmt.Printf("%s\n", *px)
}

