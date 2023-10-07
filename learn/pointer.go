package learn

import "fmt"

func zeroValue(value int) {
	value = 0
}

func zeroPointer(p *int) {
	*p = 0
}

func StartPointer() {
	fmt.Println("----------------------- Start Pointer -----------------------")
	i := 1
	fmt.Println("i =", i)

	zeroValue(i)
	fmt.Println("i =", i)

	zeroPointer(&i)
	fmt.Println("i =", i)
	fmt.Println("----------------------- End Pointer -----------------------")
}
