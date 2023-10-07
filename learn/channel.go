package learn

import (
	"fmt"
	"time"
)

func counting01() {
	for i := 0; i < 10; i++ {
		fmt.Println("counting (1): ", i)
	}
}

func counting02() {
	for i := 0; i < 10; i++ {
		fmt.Println("counting (2): ", i)
	}
}

func process(c chan string, data string) {
	c <- data
}

func StartChannel() {
	//gorutine
	go counting01()
	go counting02()
	time.Sleep(1 * time.Second)

	ch := make(chan string)
	go process(ch, "role_01")
	fmt.Println(<-ch)
	time.Sleep(1 * time.Second)

}
