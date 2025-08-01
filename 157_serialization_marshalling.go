package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)


type User struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

func main(){
	user := User{Name:"Alice", Email: "alice@example.com"}
	fmt.Println(user)
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))

	var user1 User
	err1 := json.Unmarshal(jsonData, &user1)
	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Println("User created from json data", user1)

	// marshal and unmarshal for quickly serializing or deserializing data in memory.

	jsonData1 := `{"name":"John", "email":"john@example.com"}`
	reader := strings.NewReader(jsonData1)
	decoder := json.NewDecoder(reader)

	var user2 User
	err2 := decoder.Decode(&user2)
		if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(user2)

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)

	err3 :=encoder.Encode(user)
		if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Println("Encoded json string:", buf.String())
}

// marshall and unmarshall better for quick in memory tasks
// decoder and encoder good for large data, such as API data or for streaming data.


