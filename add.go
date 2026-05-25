package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func sumOfPrimesNumbers(number int) int {
	//var primes []int
		sum := 0

	for i := 2; i < number; i++ {
		isPrime := true

		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}

		// if isPrime {
		// 	primes = append(primes, i)
		// }

		if isPrime{
			sum+=i
		}
	}

	// for _, num := range primes {
	// 	sum += num
	// }
	return sum

}

func isPrime(n int) bool{
	if n < 2{
		return false
	}

	for i := 2; i*i < n; i++{
		if n%i == 0{
			return false
		}
	}
	return true
}



func main() {
	//fmt.Println(isPrime(3))
	
	if len(os.Args) < 2{
		fmt.Println("Usage: go run <number>")
		return
	}
	input := os.Args[1]
	num, err := strconv.Atoi(input)
	if err != nil{
		fmt.Println("error: converting input")
	}
	fmt.Println(sumOfPrimesNumbers(num))



}
