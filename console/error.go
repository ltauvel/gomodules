package console

import (
	"log"
)

func printError(text string) {
	log.Fatal("\x1b[31;1m" + text + "\033[0m")
}