package learn

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	ID   int
	Name string
}

func StartDemoJson() {
	data, _ := json.Marshal(&employee{1, "test"})
	fmt.Println(string(data))

	e := employee{}
	err := json.Unmarshal([]byte(`{"ID":101,"Name":"test"}`), &e) // Corrected field name to "Name"
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(e)
}
