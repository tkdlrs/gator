package main

import (
	"fmt"

	"github.com/tkdlrs/gator/internal/config"
)

func main() {
	fmt.Println("start")
	fmt.Println("----------------------------")
	config.Read()
	fmt.Println("----------------------------")
	fmt.Println("----------------------------")
	fmt.Println("end")
}
