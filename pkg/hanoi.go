package pkg

import "fmt"

func TowersOfHanoi(n int, frompeg string, topeg string, auxpeg string) {
	if n == 1 {
		fmt.Printf("Move disk 1 from peg %s to peg %s\n", frompeg, topeg)
		return
	}

	TowersOfHanoi(n-1, frompeg, auxpeg, topeg)

	fmt.Printf("Move disk %d from peg %s to peg %s\n", n, frompeg, topeg)

	TowersOfHanoi(n-1, auxpeg, topeg, frompeg)
}
