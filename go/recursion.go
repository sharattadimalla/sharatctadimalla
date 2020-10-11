// recursion example

package main

import (
	"fmt"
	)

// define a factorial func

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)

}

func main() {
	fmt.Println(fact(5))
	
}
