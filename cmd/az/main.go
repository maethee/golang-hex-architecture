package main

import (
	"fmt"
	"os"
)

func main() {

	switch os.Args[1] {
	case "psp":
		fmt.Printf("PSP Cmd... %t", os.Args[2:])
		ExecPSPCmd(os.Args[2:])
	case "myaxa":
		fmt.Printf("Myaxa Cmd... %t", os.Args[2:])
		ExecMyaxaCmd(os.Args[2:])
	case "product":
		fmt.Println("Product Cmd...")
	}
}
