package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Find two sided primes !")
	router := mux.NewRouter()
	twoSidedPrimeRoutes(router)
	log.Fatal(http.ListenAndServe(":8001", router))
}

func twoSidedPrimeRoutes(router *mux.Router) {
	router.HandleFunc("/", helloMessage)
	prefix := "/twosidedprime"
	router.HandleFunc(prefix+"/{id}", twoSidedPrime).Methods("GET")
}

func twoSidedPrime(writer http.ResponseWriter, request *http.Request) {
	b := isTwoSidedPrime(mux.Vars(request)["id"])
	writer.Write([]byte(fmt.Sprintf("%t",b)))
}
func helloMessage(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello Two sided Prime !!!! Please hit /twosidedprime/<number>"))
}

// isTwoSidedPrime verifies whether the number passed is two sided prime or not
func isTwoSidedPrime(inputStr string) bool {
	numberOfDigits := len(inputStr)
	if checkInputContainsZero(inputStr) {
		return false
	}
	input, err := strconv.Atoi(inputStr)
	if err != nil {
		log.Fatal(err.Error())
	}
	allPrimesArr := SieveOfEratosthenes(input)
	isLeftPrime := leftThruPrime(input, numberOfDigits, allPrimesArr)
	isRightPrime := rightThruPrime(input, numberOfDigits, allPrimesArr)
	return isLeftPrime && isRightPrime
}

// func to check whether the numeric value represented in string contains 0
func checkInputContainsZero(input string) bool {
	if strings.Contains(input, "0") {
		return true
	} else {
		return false
	}
}

// function to get all prime numbers less than or equal to n
func SieveOfEratosthenes(n int) map[int]bool {
	// Create a boolean array "prime[0..n]" and initialize
	// all entries it as true. A value in prime[i] will
	// finally be false if i is Not a prime, else true.
	integers := make([]bool, n+1)
	for i := 2; i < n+1; i++ {
		integers[i] = true
	}

	for p := 2; p*p <= n; p++ {
		// If integers[p] is not changed, then it is a prime
		if integers[p] == true {
			// Update all multiples of p
			for i := p * 2; i <= n; i += p {
				integers[i] = false
			}
		}
	}

	// return all prime numbers <= n
	var primes = make(map[int]bool)
	for p := 2; p <= n; p++ {
		if integers[p] == true {
			primes[p] = true
		}
	}

	return primes
}

func leftThruPrime(n int, len int, allPrimes map[int]bool) bool {
	var mod int
	for i := len; i > 0; i-- {
		mod = power(10, i)
		if !allPrimes[n%mod] {
			return false
		}
	}
	return true
}

func rightThruPrime(n int, len int, allPrimes map[int]bool) bool {
	for n >= 1 {
		if !allPrimes[n] {
			return false
		}
		n = n/10
	}
	return true
}

func power(x int, y int) int {
	if y == 0 {
		return 1
	} else if y%2 == 0 {
		return power(x, y/2) * power(x, y/2)
	} else {
		return x * power(x, y/2) * power(x, y/2)
	}
}
