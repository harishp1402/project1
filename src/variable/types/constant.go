package constants

import (
	"fmt"
)

func Alpha() {
	var a int = 36
	b := 30.0
	fmt.Printf("a is of type %T\n", a)
	fmt.Printf("b is of type %T\n", b)

	var x, y, z = 8, 9, "Hello"
	fmt.Printf("x,y,z is of type %T, %T, %T\n", x, y, z)
}
