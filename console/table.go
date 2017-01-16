package console

import (
	"os"
	"fmt"
	"strconv"
	"strings"
	"text/tabwriter"
)

func printTable(header []string, content [][]string, indent int) {
	
	// Generate table sperators line
		
		var separator = []string{}
		
		// Generate sperator based on header length
		for _, h := range header {
			separator = append(separator, strings.Repeat("-", len(h)))
		}
		
		// Generate sperator based on content length
		for _, c := range content {
			for i := 0; i < len(c); i++ {
				contentlen := len(c[i])

				if i > len(separator) - 1 {
					separator = append(separator, strings.Repeat("-", contentlen))
				} else {
					if contentlen > len(separator[i]) {
						separator[i] = strings.Repeat("-", contentlen)
					}
				}
			}
		}
	
	// Generate a table containing headers, sperators, content and footer
	
		table := [][]string{ header }
		table = append(table, separator)
		table = append(table, content...)
		table = append(table, separator)
		
	// Print the table
	
		fmt.Println("")
		tab := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)	
		for _, r := range table {
			fmt.Fprintln(tab, strings.Repeat(" ", indent), strings.Join(r, " \t"))
		}
		tab.Flush()
		
		fmt.Println(strings.Repeat(" ", indent), "TOTAL: " + strconv.Itoa(len(content)))
		fmt.Println("")		

}