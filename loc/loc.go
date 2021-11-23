package main

import (
	"fmt"
	"os"
)

func main() {
	where, _ := os.Executable()
	fmt.Printf("Location: %s\n", where)
}
