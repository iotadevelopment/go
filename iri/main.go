package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}
	//port, _ := strconv.Atoi(arguments[1])

	fmt.Println("  _____ _____  _____   ___    ___")
	fmt.Println(" |_   _|  __ \\|_   _| |__ \\  / _ \\")
	fmt.Println("   | | | |__) | | |      ) || | | |")
	fmt.Println("   | | |  _  /  | |     / / | | | |")
	fmt.Println("  _| |_| | \\ \\ _| |_   / /_ | |_| |")
	fmt.Println(" |_____|_|  \\_\\_____| |____(_)___/")
	fmt.Println("")
	fmt.Println("Started successfully ...")
	fmt.Println("")

}
