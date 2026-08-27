package main

import (
	"fmt"

	"devpath.local/go0101/internal/store"
)

func main() {
	fmt.Println("same module can import internal:", store.Kind())
}
