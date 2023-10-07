package learn

import "fmt"

func StartLoop() {
	fmt.Println("----------------------- Start Loop -----------------------")
	list_number := []int{1, 2, 3, 4, 5}
	size := len(list_number)
	for count := 0; count < size; count++ {
		fmt.Println(list_number[count])
	}

	for index, value := range list_number {
		fmt.Println("Index = ", index, "value = ", value)
	}

	for _, value := range list_number {
		fmt.Println("value = ", value)
	}
	fmt.Println("----------------------- End Loop -----------------------")
}
