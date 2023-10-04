package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	do_token := os.Getenv("DIGITALOCEAN_TOKEN")
	masterkey, ok := os.LookupEnv("MASTERKEY")
	if !ok {
		masterkey = "secret2key"
	}
	fmt.Println("do_token:", do_token, "masterkey:", masterkey)
}
