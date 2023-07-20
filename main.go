package main

import "recursion/pkg"

func main() {
	// var n int
	// fmt.Print("Enter number : ")
	// fmt.Scanf("%d", &n)

	// fmt.Println("factorial of", n, ":", fact(n))
	// fmt.Println(RecursivePrinter(5))
	// n := 14
	// for i := 0; i <= n; i++ {
	// 	fmt.Println(pkg.Fibonacci(i))
	// }

	pkg.TowersOfHanoi(5, "A", "C", "B")
}
