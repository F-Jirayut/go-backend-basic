package learn

import "fmt"

func StartArray() {
	fmt.Println("----------------------- Start Array -----------------------")
	// array
	var list = [5]int{}
	list[0] = 100
	list[1] = 214
	list[2] = 34
	list[3] = 46
	fmt.Println(list)
	fmt.Println(list[0])
	var items = [3]string{"a", "b", "c"}
	fmt.Println(items)

	// slice
	var listSlice = []int{}
	listSlice = append(listSlice, 5, 4, 6)
	fmt.Println(listSlice)
	fmt.Println("----------------------- End Array -----------------------")
}
