package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	nums := make([]int, 3)

	for {
		inp, err := getInputStdIn()

		if err != nil {
			panic(err)
		}

		if isExit(inp) {
			fmt.Println("[INFO] Entered 'x'. Exit.")
			break
		}

		num, err := isInt(inp)
		if err != nil {
			fmt.Println("[ERROR] Type error. Enter integer number")
			continue
		}

		nums = append(nums, num)
		sort.Ints(nums)
		fmt.Println(nums)
	}
}

func isInt(str string) (int, error) {
	return strconv.Atoi(str)
}

func isExit(str string) bool {
	return "x" == strings.ToLower(str)
}

func getInputStdIn() (string, error) {
	fmt.Println("[INFO] Enter integer number: ")
	inp, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(inp, "\n"), nil
}
