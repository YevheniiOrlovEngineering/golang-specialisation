package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const ArrSizeMax = 10
	var nums []int

	nums = fillArray(nums, ArrSizeMax)
	bubbleSort(nums)
	printArray(nums)
}

func fillArray(nums []int, ArrSizeMax int) []int {
	for len(nums) < ArrSizeMax {
		nums = append(nums, getElem())
		fmt.Printf("Entered numbers: %d, Remain: %d\n", len(nums), ArrSizeMax-len(nums))
	}
	return nums
}

func getElem() int {
	for {
		inp := readStdIn()
		num, err := strToInt(inp)
		if err != nil {
			fmt.Println("[ERROR] Type error. Enter integer number")
			continue
		}
		return num
	}
}

func readStdIn() string {
	fmt.Println("[INFO] Enter integer number: ")
	inp, err := bufio.NewReader(os.Stdin).ReadString('\n')
	isErr(err)
	return strings.TrimSuffix(inp, "\n")
}

func strToInt(inp string) (int, error) {
	return strconv.Atoi(inp)
}

func bubbleSort(nums []int) {
	for r := len(nums) - 1; r > 0; r-- {
		for i := 0; i < r; i++ {
			if nums[i] > nums[i+1] {
				swap(nums, i)
			}
		}
	}
}

func swap(s []int, i int) {
	tmp := s[i]
	s[i] = s[i+1]
	s[i+1] = tmp
}

func printArray(nums []int) {
	fmt.Print("Sorted Array: ")
	for _, num := range nums {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}

func isErr(e error) {
	if e != nil {
		panic(e)
	}
}
