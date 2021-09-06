package main

import (
	"log"
	"os"

	"github.com/mxschmitt/playwright-go"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("USAGE: go-playwright-installer <driver-directory> [browsers]")
	}

	driverDir := os.Args[1]

	var browsers []string
	if len(os.Args) > 2 {
		browsers = append(browsers, os.Args[2:]...)
	}

	if err := playwright.Install(&playwright.RunOptions{
		DriverDirectory: driverDir,
		Browsers:        browsers,
	}); err != nil {
		log.Fatal(err)
	}
}
