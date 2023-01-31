package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	scrapCmd := flag.NewFlagSet("scrap", flag.ExitOnError)
	code := scrapCmd.String("code", "", "code")

	if len(os.Args) < 2 {
		fmt.Println("expected subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "scrap":
		scrapCmd.Parse(os.Args[2:])
		ExampleScrape(*code)
	default:
		fmt.Println("expected subcommands.")
		os.Exit(1)
	}
}
