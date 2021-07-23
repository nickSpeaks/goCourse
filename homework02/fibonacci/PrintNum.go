package fibonacci

import "fmt"

func PrintNum(n int) {
	n1, n2, n3 := 0, 1, 1
	fmt.Printf("%d %d ", n1, n2)

	for i := 2; i < n; i++ {
		n3 = n1 + n2
		fmt.Printf("%d ", n3)
		n1 = n2
		n2 = n3
	}
}
