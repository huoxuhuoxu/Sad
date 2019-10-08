package main

import (
	"log"
	"os"
)

func ToError(errMsg string) {
	log.Println(errMsg)
	os.Exit(1)
}

func main() {
	args := os.Args
	if len(args) < 3 {
		ToError("Error, 001")
	}

	switch args[1] {
	case "run":
		log.Println("准备启动", args[2])

	default:
		ToError("Error, 002")
	}

}
