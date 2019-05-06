// +build ignore

package main

import (
	"fmt"
	"log"
	"time"
	"math"
	"math/rand"
	"math/big"
	crand "crypto/rand"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("rand: ", rand.Intn(10))

	seed, err := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))

	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(seed.Int64())
	fmt.Println("rand: ", rand.Intn(10))
}