package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// declaring a struct
type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

func main() {
	fmt.Println("HI")
	file, _ := ioutil.ReadFile("./users.json")

	data := Users{}

	_ = json.Unmarshal([]byte(file), &data)

	for i := 0; i < len(data.Users); i++ {
		fmt.Println("Name: ", data.Users[i].Name)
		fmt.Println("Gender: ", data.Users[i].Gender)
	}

}
