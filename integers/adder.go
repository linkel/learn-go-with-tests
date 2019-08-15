package integers

import "fmt"

// Add will add two integers together, returns the sum.
func Add(x, y int) int {
	return x + y
}

// ExampleAdd is an example. Can use examples to assist with documentation like so.
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output : 6
}
