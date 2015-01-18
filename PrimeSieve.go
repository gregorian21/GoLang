package main

import (
	"fmt"
	"time"
	"log"
)

const (
	limit = 100000 //size of the slice
)

func timeTrack(start time.Time, name string) { //ref: https://coderwall.com/p/cp5fya/measuring-execution-time-in-go
    elapsed := time.Since(start)
    log.Printf("%s took %s", name, elapsed)
}

func main() {
	
	defer timeTrack(time.Now(), "Calculating primes") //time tracking
	
	nums := make([]uint32, limit) //construct a slice
	nums[0] = 1 //set 3 as a prime number
	
	for i := 1; i < len(nums); i++ { //iterate through the slice
		nums[i] = 1 //set the number as a potential prime
		for j := 0; j < i; j++ { //iterate through earlier indices
			if nums[j] == 1 && (2*i+3)%(2*j+3) == 0 { //if the earlier index was a prime and is a factor
				nums[i] = 0 //set the number to not be a prime
				break //exit the loop early
			}
		}
	}
	
	fmt.Printf("Primes:\n2\n") // 2 is a prime
	for i := 0; i < len(nums); i++ { //iterate through the slice
		if nums[i] == 1 { //if the index is a prime
			fmt.Printf("%d\n", (2*i+3)) //print the number; As we are omitting all even numbers, convert the index to the actual number
		}
	}
}

