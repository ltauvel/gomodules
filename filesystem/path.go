package filesystem

import(
	"os"
	"regexp"
	"strings"
    "path/filepath"
	"github.com/ltauvel/gomodules/console"
)

func checkPathExists(path string) bool {
	console.PrintDebug("Checking if", path, "exists")
    _, err := os.Stat(path)
    if err == nil { return true }
    if os.IsNotExist(err) { return false }
    return true
}

func joinPath(part ...string) string {
	var result string
	
	for _, p := range part {
		if len(p) > 0 {
			result = result + p + "/"
		}
	}
	
	// Return the path
	return filepath.Clean(result)
}

func getValidPath(path ...string) string {
	var result string
	
	for _, p := range path {
		if len(p) > 0 {
		
			// Check if the current path is relative
			// If so, Join with the working directory
			if match, _ := regexp.MatchString("^\\.", p); match == true {
				wd, _ := os.Getwd()
				p = filepath.Join(strings.Replace(wd, " ", "\\ ", -1), p)
			}
			
			// Check if the current path contains environment variables
			// If so, Replace with the environment variables values
			r, _ := regexp.Compile("(%[^%]+%)")
			if r.MatchString(p) {
				for _, m := range r.FindAllString(p, -1) {
					p = strings.Replace(p, m, os.Getenv(strings.Replace(m, "%", "",-1)), -1)
				}
			}
			
			// Clean the path before check exists
			p = filepath.Clean(p)
			
			// Check if the current path exist
			// If so, break and return it.
			if checkPathExists(p) {
				result = filepath.Clean(p)
				break
			}
		}
	}
	
	// Return the path
	return result
}