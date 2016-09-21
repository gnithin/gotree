/*
Utils file for common stuff
*/
package tree

import (
	"fmt"
)

/*
Helper function that's used to print out stuff when
debug is true
*/
func debug(ip ...interface{}) {
	if DEBUG {
		fmt.Println(ip...)
	}
}
