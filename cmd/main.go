package main

import (
	"fmt"
	a "zsandibe/internal/app"
)

func main() {
	if err := a.Start(); err != nil {
		fmt.Println("Server is busy...")
	}
}
