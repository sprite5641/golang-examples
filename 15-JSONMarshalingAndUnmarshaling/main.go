package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"`
}

func main() {
	person := Person{Name: "Alice", Age: 30}
	jsonData, _ := json.Marshal(person)
	fmt.Println("JSON:", string(jsonData))

	jsonStr := `{"name":"Bob","age":25}`
	var newPerson Person
	_ = json.Unmarshal([]byte(jsonStr), &newPerson)
	fmt.Println("Unmarshaled Person:", newPerson)
}
