package main

import (
	"fmt"
	"recursion/pkg"
)

func main() {
	var n int
	fmt.Print("Enter number : ")
	fmt.Scanf("%d", &n)

	fmt.Println("factorial of", n, ":", pkg.Factorial(n))
	for i := 0; i <= n; i++ {
		fmt.Println(pkg.Fibonacci(i))
	}

	pkg.TowersOfHanoi(10, "A", "C", "B")
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Is array sorted? ", pkg.IsArraySorted(array, uint(len(array))))
}
