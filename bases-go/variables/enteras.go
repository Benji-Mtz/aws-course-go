package variables

import "fmt"

func MuestroEnteros() {
	// Variable declarativa
	var intcomun int

	// variable por asignacion
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Printf("intcomun = %v \n", intcomun)
	fmt.Printf("intde32 = %v \n", intde32)
	fmt.Printf("intde64 = %v \n", intde64)
}
