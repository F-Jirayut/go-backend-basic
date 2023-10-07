package learn

import "fmt"

func StartMaps() {
	fmt.Println("----------------------- Start Map -----------------------")
	var departments = map[string]string{}
	// add or update
	departments["HR"] = "Human Resources"
	departments["IT"] = "Information Technology"
	departments["SALES"] = "Sales Department"
	fmt.Println(departments)
	fmt.Println(departments["HR"])

	var numbers = map[int]int{0: 150, 1: 100}
	numbers[10] = 1000
	// delete
	delete(numbers, 10)
	fmt.Println(numbers)
	fmt.Println("----------------------- End Map -----------------------")
}
