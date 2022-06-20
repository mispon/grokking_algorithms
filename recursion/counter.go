package recursion

import "fmt"

func counter(n int) {
	if n == 0 {
		fmt.Println("finish!")
		return
	}

	fmt.Println("counter:", n)
	counter(n - 1)
}
