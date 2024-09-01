package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	//Gets the file name from terminal and error manegement for invalid input

	if len(os.Args) < 2 {
		fmt.Println("Example input : go run your-program.go data.txt")
	} else if len(os.Args) > 2 {
		fmt.Println("Example input : go run your-program.go data.txt")
	}

	fileName := os.Args[1]

	//Read file content
	allNums, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	//Slice that contains numbers
	var nums []float64

	//Converts string to integer(Atoi), adds it to a slice and remove unwanted chars
	numStrings := strings.Split(strings.TrimSpace(string(allNums)), "\n")
	for i := 0; i < len(numStrings); i++ {
		num, _ := strconv.Atoi(numStrings[i])
		x := float64(num)
		nums = append(nums, x)
	}

	// Using different functions for different calculations
	a := avrage(nums)
	m := median(nums)
	v := variance(nums)
	sd := sDeviation(v)

	//Print out the results
	fmt.Println("Average:", a)
	fmt.Println("Median:", m)
	fmt.Println("Variance:", int(math.Round(v)))
	fmt.Println("Standard Deviation:", sd)
}

func avrage(nums []float64) int {
	x := 0.0
	for _, num := range nums {
		x += num
	}
	rounded := x / float64(len(nums))
	rounded = math.Round(rounded)
	return int(rounded)
}

func median(nums []float64) int {

	//sort in asecending order
	sort.Float64s(nums)

	n := len(nums)

	//Determine wheter the length of the slice is odd or even
	isOdd := n%2 != 0

	var median float64

	if isOdd {
		//if the lenght is odd, the median is the middle element
		median = float64(nums[n/2])
	} else {
		// If the lenght is even, the median is the avarage of the two middle elements
		median = float64((nums[n/2-1] + nums[n/2]) / 2)
	}

	median = math.Round(median)

	return int(median)
}

func variance(nums []float64) float64 {
	var x float64
	var xMean float64 //avrage of all the numbers
	var squaredDifferences float64

	for _, num := range nums {
		x += num
	}

	xMean = x / float64(len(nums))

	for _, value := range nums {
		squaredDifferences += math.Pow(value-xMean, 2)
	}

	s2 := squaredDifferences / float64(len(nums))

	return s2
}

func sDeviation(v float64) int {

	s := math.Sqrt(v)
	s = math.Round(s)
	return int(s)
}
