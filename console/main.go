package console


// Output prints
func PrintSection(text string, args ...string) {
	printSection(text, args...)
}

func Print(text string, args ...string) {
	print(text, args...)
}


// Debug  prints
var Debug bool

func PrintDebugSection(text string) {
	printDebugSection(text) 
}

func PrintDebug(text string, args ...string) {
	printDebug(text, args...)
}


// Error  prints
func PrintError(text string) {
	printError(text)
}

// Table  prints
func PrintTable(header []string, content [][]string, indent int) {
	printTable(header, content, indent)
}