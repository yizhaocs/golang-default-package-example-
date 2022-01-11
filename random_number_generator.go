package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	generatesRandomIntegerGivenBetweenZeroAndMaxValue()
	generatesRandomIntegerGivenBetweenZeroAndMaxValue()
	generatesRandomIntegerGivenBetweenZeroAndMaxValue()
	generateRrandomNumberBetweenTwoValues()
	generateRrandomNumberBetweenTwoValues()
	generateRrandomNumberBetweenTwoValues()
	realRandomNumber()
	realRandomNumber()
	realRandomNumber()
}

// generates a random integer given between zero and a max value
func generatesRandomIntegerGivenBetweenZeroAndMaxValue(){
	fmt.Println("******************** generatesRandomIntegerGivenBetweenZeroAndMaxValue ********************")
	n := 10
	for i := 0; i < n; i++ {
		fmt.Println(rand.Intn(n))
	}
}

// generate a random number between two values
func generateRrandomNumberBetweenTwoValues(){
	fmt.Println("******************** generateRrandomNumberBetweenTwoValues ********************")
	min := 10
	max := 30
	fmt.Println(rand.Intn(max - min) + min)
}

/*
	Indeed on the official documentation we can read

	Top-level functions, such as Float64 and Int,
	use a default shared Source that produces a deterministic sequence of values each time a program is run.
	Use the Seed function to initialize the default Source if different behavior is required for each run.
 */
func realRandomNumber(){
	fmt.Println("******************** realRandomNumber ********************")
	/*
		So we're using "seed := time.Now().UnixNano()" to generate a different 64 number.
		因为"func (t Time) UnixNano() int64"的返回值是int64
		Every single time we start up our program, we use that as the seed to generate a new source object,
	*/
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	min := 10
	max := 30
	fmt.Println(rand.Intn(max - min + 1) + min)
}