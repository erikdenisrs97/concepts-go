package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   uint   `json:"age"`
}

func main() {
	// json.Marshal() struct to json
	// json.Unmarshal() json to struct
	d := dog{"Bya", "Shitzu", 4}
	fmt.Println(d)

	dJSON, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dJSON) // slice of bytes

	fmt.Println(bytes.NewBuffer(dJSON))

	dJSON2 := `{"name":"Bya","breed":"Shitzu","age":4}`

	var d2 dog

	if err := json.Unmarshal([]byte(dJSON2), &d2); err != nil {
		log.Fatal(err)
	}

	fmt.Println(d2)
}
