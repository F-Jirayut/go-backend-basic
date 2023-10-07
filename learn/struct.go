package learn

import "fmt"

type User struct {
	id     int
	name   string
	age    int
	active bool
}

func StartStruct() {
	fmt.Println("----------------------- Start Struct -----------------------")
	user := User{
		id:     1,
		name:   "test",
		age:    20,
		active: true,
	}
	fmt.Println(user)

	users := []User{}
	users = append(users, User{
		id:     1,
		name:   "test",
		age:    20,
		active: true,
	})
	fmt.Println(users)
	fmt.Println("----------------------- End Struct -----------------------")
}
