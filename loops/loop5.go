//loops: utilizando ASCII*
//%#U = unicodego ru
//

package main

import "fmt"

func main() {

	for i := 33; i <= 122; i++ {
		fmt.Printf("%d - %v\n", i, string(i))
	}
}
