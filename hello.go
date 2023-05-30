package main

import (
	"fmt"

	"github.com/gravitl/netmaker/models"
)

func main() {
	user := models.User{
		UserName: "matt",
		Password: "kasun",
		IsAdmin:  true,
	}
	fmt.Println("hello world!")
	fmt.Println("another message")
	fmt.Println("third message")
	fmt.Println(user)
}
