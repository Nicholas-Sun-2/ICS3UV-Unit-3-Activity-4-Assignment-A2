/**
 * @author Nicholas Sun
 * @version 1.0.0
 * @date 2025-11-21
 * @fileoverview This program tests if you are in the correct weight range to enter the pie-eating contest.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Constants
	const MIN_WEIGHT float64 = 77.0
	const MAX_WEIGHT float64 = 105.0
	var weight float64

	// Declaring variables
	var weightAsString string

	// Input
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("How much do you weigh (kg)? ")
	weightAsString, _ = reader.ReadString('\n')
	weightAsString = strings.TrimSpace(weightAsString)
	// ParseFloat usage from https://www.includehelp.com/golang/strconv-parsefloat-function-with-examples.aspx
	weight, _ = strconv.ParseFloat(weightAsString, 64) // This is the only way I found to convert strings into floats.

	// If statement
	if weight >= MIN_WEIGHT && weight <= MAX_WEIGHT {
		fmt.Println("You may enter the contest.")
	} else {
		fmt.Println("You may not enter the contest.")
	}
}
