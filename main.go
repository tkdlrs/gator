package main

import (
	"fmt"

	"github.com/tkdlrs/gator/internal/config"
)

func main() {
	fmt.Println("start")
	fmt.Println("----------------------------")
	//
	fmt.Println("Start first Read")
	aConfig, err := config.Read()
	if err != nil {
		fmt.Println("Error in first read")
		return
	}
	fmt.Println("Fin first Read")
	//
	fmt.Println("Start set user")
	aConfig.SetUser()
	fmt.Println("Fin set user")
	//
	fmt.Println("Start second Read")
	uConfig, err := config.Read()
	if err != nil {
		fmt.Println("Error in second read")
		return
	}
	fmt.Println("Fin second Read")
	fmt.Println(uConfig)
	//
	fmt.Println("----------------------------")
	fmt.Println("end")
}
