package learn

import "fmt"

// Define the car interface
type car interface {
	IsRunning() bool
}

// Define the status struct
type status struct {
	run bool
}

// Implement the IsRunning method for the status struct
func (s status) IsRunning() bool {
	return s.run
}

// CarIsRunning takes a car interface and checks if it's running
func CarIsRunning(c car) {
	fmt.Println(c.IsRunning())
}

func StartInterface() {
	fmt.Println("----------------------- Start Interface -----------------------")
	car1 := status{run: false}
	car2 := status{run: true}

	CarIsRunning(car1)
	CarIsRunning(car2)
	fmt.Println("----------------------- Start Interface -----------------------")
}
