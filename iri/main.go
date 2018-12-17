package main

import (
	"encoding/base64"
	"fmt"
	"github.com/iotadevelopment/go/ternary"
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

	transactionBytes, _ := base64.StdEncoding.DecodeString("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")
	trits := ternary.BytesToTrits(sDec)[:8019]
}
