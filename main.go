package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1]
	if len(os.Args) != 1 {
		fmt.Println("Message error !")
		return
	}
	
}
