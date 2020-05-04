package main

import (
	"flag"
	"fmt"

	"github.axa.com/maethee-chakkuchantorn/psp-improvement/internal/app/uc"
)

type CommandHandler struct {
	ucHandler uc.Handler
}

func initPSPCommand() CommandHandler {
	ucInterface := uc.HandlerConstructor{}.New()
	return CommandHandler{
		ucHandler: ucInterface,
	}
}

func ExecPSPCmd(args []string) {
	cmd := initPSPCommand()

	formatCmd := flag.NewFlagSet("format", flag.ExitOnError)
	formatType := formatCmd.String("type", "TCP", "format pattern")
	formatInput := formatCmd.String("input", "", "input file path")
	formatOutput := formatCmd.String("output", "", "output file path")
	formatConfig := formatCmd.String("config", "", "config file path")

	switch args[0] {
	case "format":
		formatCmd.Parse(args[1:])
		fmt.Println("Input: ", *formatInput)
		fmt.Println("Output: ", *formatOutput)
		fmt.Println("Type: ", *formatType)
		fmt.Println("Config: ", *formatConfig)

		config := uc.LoadConfiguration(*formatConfig)

		fmt.Println(config)

		if *formatType == "TCP" {
			fmt.Println("TCP format proceed: ")
			cmd.ucHandler.FormatTCP(*formatInput, *formatOutput, config)
		}
	}
}
