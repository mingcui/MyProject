// 4.9-3
package main

import "fmt"

func math(num1, num2 int) (sum, account, div int) {
	sum = num1 + num2
	account = num1 * num2
	div = num2 / num1
	return sum, account, div
}
func main() {
	num1, num2 := 2, 4
	sum, account, div := math(num1, num2)
	fmt.Printf("%d Add %d Is %d\n ", num1, num2, sum)
	fmt.Printf("%d Account %d Is %d\n", num1, num2, account)
	fmt.Printf("%d div %d Is %d\n", num2, num1, div)

}
