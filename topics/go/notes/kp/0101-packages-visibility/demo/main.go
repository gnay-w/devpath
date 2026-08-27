package main

import (
	"fmt"

	"devpath.local/go0101/user"
)

func main() {
	u := user.NewUser("gnay")
	fmt.Println("via Name():", u.Name())
}
