package console

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
)

func printSection(text string, args ...string) {
	section := color.New(color.FgYellow).Add(color.Underline)
	section.Println("\r\n" + text + " " + strings.Join(args, " ") + "\r\n* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * \r\n")
}

func print(text string, args ...string) {
	fmt.Println(text + " " + strings.Join(args, " "))
}


