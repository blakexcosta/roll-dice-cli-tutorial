package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

func main() {
	// This project is released with goreleaser

	// seed the random number generator so it doesn't give you the same seed every time running
	rand.Seed(time.Now().UTC().UnixNano())

	// kingping and cobra are more complex tools, otherwise the flag package here is fine
	// flag number 1
	dice := flag.String("d", "d6", "The type of dice to roll. Format: dx where X is an integer Default: d6")
	// flag number 2
	numRoll := flag.Int("n", 1, "The number of die to roll, Default: 1")
	flag.Parse()

	// print out chosen die
	fmt.Println("You chose a", *dice)

	matched, _ := regexp.Match("d\\d+", []byte(*dice)) // cast into byteslice, because can't match on a string but can on a byteslice
	fmt.Println(matched)

	// matching
	if matched == true {
		rolls := RollDice(dice, numRoll)
		PrintDice(rolls)
	} else {
		log.Fatalf("Improper format for dice. Format should be dX where X is an integer")
	}
}

func RollDice(dice *string, times *int) []int {
	var rolls []int
	diceSides := (*dice)[1:]
	d, err := strconv.Atoi(diceSides)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < *times; i++ {
		rolls = append(rolls, rand.Intn(d)+1)
	}
	return rolls
}

func PrintDice(rolls []int) {
	for i, dice := range rolls {
		fmt.Printf("Roll %d was %d\n", i+1, dice)
	}
}

// TODO: take highest number roll
func RollWithAdvantage() {

}

// TODO: take lowest number roll
func RollWithDisadvantage() {

}
