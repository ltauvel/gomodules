package console

import (
	"github.com/fatih/color"
	"strings"
)

func printDebugSection(text string) {
	if Debug {
		separator := strings.Repeat("*", 65)
		if len(text) > 0 {
			text = strings.Repeat("*", 6) + " " + text + " " + separator
			text = text[:len(separator)]
		} else {
			text = separator
		}
		
		printDebug("")
		printDebug(text) 
		printDebug("")
	}
}

func printDebug(text string, args ...string) {
	if Debug {
		debug := color.New(color.FgCyan).Add(color.Underline)
		debug.Println("DEBUG: " + text + " " + strings.Join(args, " "))
	}
}